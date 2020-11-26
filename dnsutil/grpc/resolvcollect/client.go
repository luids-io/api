// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package resolvcollect

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Client provides a grpc client.
type Client struct {
	opts   clientOpts
	logger yalogi.Logger
	//grpc connection
	conn   *grpc.ClientConn
	client pb.ResolvCollectClient
	//control
	closed bool
}

// ClientOption encapsules options for client.
type ClientOption func(*clientOpts)

type clientOpts struct {
	logger    yalogi.Logger
	closeConn bool
}

var defaultClientOpts = clientOpts{
	logger:    yalogi.LogNull,
	closeConn: true,
}

// CloseConnection option closes grpc connection on shutdown.
func CloseConnection(b bool) ClientOption {
	return func(o *clientOpts) {
		o.closeConn = b
	}
}

// SetLogger option allows set a custom logger.
func SetLogger(l yalogi.Logger) ClientOption {
	return func(o *clientOpts) {
		if l != nil {
			o.logger = l
		}
	}
}

// NewClient returns a new client.
func NewClient(conn *grpc.ClientConn, opt ...ClientOption) *Client {
	opts := defaultClientOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Client{
		opts:   opts,
		logger: opts.logger,
		conn:   conn,
		client: pb.NewResolvCollectClient(conn),
	}
}

// Collect implements dnsutil.ResolvCollector interface.
func (c *Client) Collect(ctx context.Context, client net.IP, name string, resolved []net.IP, cnames []string) error {
	if c.closed {
		c.logger.Warnf("client.dnsutil.resolvcollect: collect(%v,%s,%v,%v): client is closed", client, name, resolved, cnames)
		return dnsutil.ErrUnavailable
	}
	rr := make([]string, 0, len(resolved))
	for _, r := range resolved {
		rr = append(rr, r.String())
	}
	rcnames := make([]string, 0, len(cnames))
	for _, r := range cnames {
		rcnames = append(rcnames, r)
	}
	req := &pb.ResolvCollectRequest{
		ClientIp:       client.String(),
		Name:           name,
		ResolvedIps:    rr,
		ResolvedCnames: rcnames,
	}
	_, err := c.client.Collect(ctx, req)
	if err != nil {
		c.logger.Warnf("client.dnsutil.resolvcollect: collect(%v,%s,%v,%v): %v", client, name, resolved, cnames, err)
		return c.mapError(err)
	}
	return nil
}

//mapping errors
func (c *Client) mapError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.Canceled:
		return dnsutil.ErrCanceledRequest
	case codes.InvalidArgument:
		return dnsutil.ErrBadRequest
	case codes.Unimplemented:
		return dnsutil.ErrNotSupported
	case codes.Internal:
		return dnsutil.ErrInternal
	case codes.Unavailable:
		return dnsutil.ErrUnavailable
	case codes.ResourceExhausted:
		if st.Message() == dnsutil.ErrLimitDNSClientQueries.Error() {
			return dnsutil.ErrLimitDNSClientQueries
		}
		if st.Message() == dnsutil.ErrLimitResolvedNamesIP.Error() {
			return dnsutil.ErrLimitResolvedNamesIP
		}
		return dnsutil.ErrUnavailable
	default:
		return dnsutil.ErrUnavailable
	}
}

//Close closes the client.
func (c *Client) Close() error {
	if c.closed {
		return errors.New("client closed")
	}
	c.closed = true
	if c.opts.closeConn {
		return c.conn.Close()
	}
	return nil
}

// Ping checks connectivity with the api.
func (c *Client) Ping() error {
	if c.closed {
		return errors.New("client closed")
	}
	st := c.conn.GetState()
	switch st {
	case connectivity.TransientFailure:
		return fmt.Errorf("connection state: %v", st)
	case connectivity.Shutdown:
		return fmt.Errorf("connection state: %v", st)
	}
	return nil
}

//API returns API service name implemented.
func (c *Client) API() string {
	return ServiceName()
}
