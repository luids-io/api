// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package archive

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/encoding"
	"github.com/luids-io/api/dnsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Client provides a grpc client.
type Client struct {
	opts   clientOpts
	logger yalogi.Logger
	//grpc connection
	conn   *grpc.ClientConn
	client pb.ArchiveClient
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
		client: pb.NewArchiveClient(conn),
	}
}

// SaveResolv implements dnsutil.Archiver interface.
func (c *Client) SaveResolv(ctx context.Context, rd dnsutil.ResolvData) (uuid.UUID, error) {
	if c.closed {
		c.logger.Warnf("client.dnsutil.archive: saveresolv(%v): client is closed", rd.ID)
		return uuid.Nil, dnsutil.ErrUnavailable
	}
	//create request
	rdpb := &pb.ResolvData{}
	err := encoding.ResolvDataPB(&rd, rdpb)
	if err != nil {
		c.logger.Warnf("client.dnsutil.archive: saveresolv(%v): %v", rd.ID, err)
		return uuid.Nil, dnsutil.ErrBadRequest
	}
	//do save
	resp, err := c.client.SaveResolv(ctx, &pb.SaveResolvRequest{Resolv: rdpb})
	if err != nil {
		c.logger.Warnf("client.dnsutil.archive: saveresolv(%v): %v", rd.ID, err)
		return uuid.Nil, c.mapError(err)
	}
	rid, err := uuid.Parse(resp.GetId())
	if err != nil {
		c.logger.Warnf("client.dnsutil.archive: saveresolv(%v): invalid returned id: %v", rd.ID, err)
		return uuid.Nil, c.mapError(err)
	}
	return rid, nil
}

//mapping errors.
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
	default:
		return dnsutil.ErrUnavailable
	}
}

//Close closes the client
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

// Ping checks connectivity with the api
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

//API returns API service name implemented
func (c *Client) API() string {
	return ServiceName()
}
