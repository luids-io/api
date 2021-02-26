// Copyright 2021 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package finder

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

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
	client pb.FinderClient
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
		client: pb.NewFinderClient(conn),
	}
}

// GetResolv implements dnsutil.Finder interface
func (c *Client) GetResolv(ctx context.Context, id string) (dnsutil.ResolvData, bool, error) {
	if c.closed {
		c.logger.Warnf("client.dnsutil.finder: getresolv(%s): client is closed", id)
		return dnsutil.ResolvData{}, false, dnsutil.ErrUnavailable
	}
	//check req
	if id == "" {
		c.logger.Warnf("client.dnsutil.finder: getresolv(): id is empty")
		return dnsutil.ResolvData{}, false, dnsutil.ErrBadRequest
	}
	//do req
	resp, err := c.client.GetResolv(ctx, &pb.GetResolvRequest{Id: id})
	if errNotFound(err) {
		return dnsutil.ResolvData{}, false, nil
	}
	if err != nil {
		c.logger.Warnf("client.dnsutil.finder: getresolv(%s): %v", id, err)
		return dnsutil.ResolvData{}, false, c.mapError(err)
	}
	if resp.GetData() == nil {
		c.logger.Errorf("client.dnsutil.finder: getresolv(%s): unexpected data empty", id)
		return dnsutil.ResolvData{}, false, dnsutil.ErrInternal
	}
	//map data to object
	data, err := encoding.ResolvData(resp.GetData())
	if err != nil {
		c.logger.Warnf("client.dnsutil.finder: getresolv(%s): converting data: %v", id, err)
		return dnsutil.ResolvData{}, false, dnsutil.ErrInternal
	}
	return data, true, nil
}

// ListResolvs implements dnsutil.Finder interface
func (c *Client) ListResolvs(ctx context.Context, filters []dnsutil.ResolvsFilter,
	rev bool, max int, next string) ([]dnsutil.ResolvData, string, error) {
	if c.closed {
		c.logger.Warnf("client.dnsutil.finder: listresolvs(): client is closed")
		return nil, "", dnsutil.ErrUnavailable
	}
	if max < 0 {
		c.logger.Warnf("client.dnsutil.finder: listresolvs(): invalid max")
		return nil, "", dnsutil.ErrBadRequest
	}
	//create request
	req := &pb.ListResolvsRequest{
		Max:     int32(max),
		Next:    next,
		Reverse: rev,
		Filters: make([]*pb.ResolvsFilter, 0, len(filters)),
	}
	for _, f := range filters {
		filterpb, err := encoding.ResolvsFilterPB(f)
		if err != nil {
			c.logger.Warnf("client.dnsutil.finder: listresolvs(): bad filter: %v", err)
			return nil, "", dnsutil.ErrBadRequest
		}
		req.Filters = append(req.Filters, filterpb)
	}
	//do list
	resp, err := c.client.ListResolvs(ctx, req)
	if err != nil {
		c.logger.Warnf("client.dnsutil.finder: listresolvs(%s): %v", err)
		return nil, "", c.mapError(err)
	}
	//process response
	data := make([]dnsutil.ResolvData, 0, len(resp.GetData()))
	for _, rpb := range resp.GetData() {
		r, err := encoding.ResolvData(rpb)
		if err != nil {
			c.logger.Errorf("client.dnsutil.finder: listresolvs(%s): encoding returned data: %v", err)
			return nil, "", c.mapError(dnsutil.ErrInternal)
		}
		data = append(data, r)
	}
	return data, resp.GetNext(), nil
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

func errNotFound(err error) bool {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	if !ok {
		return false
	}
	if st.Code() == codes.NotFound {
		return true
	}
	return false
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
