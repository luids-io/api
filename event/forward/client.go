// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package forward

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/event/encoding"
	pb "github.com/luids-io/api/protogen/eventpb"
	"github.com/luids-io/core/event"
	"github.com/luids-io/core/utils/yalogi"
)

// Client is the main struct for grpc client
type Client struct {
	opts   clientOpts
	logger yalogi.Logger
	//grpc connection
	conn   *grpc.ClientConn
	client pb.ForwardClient
	//control
	closed bool
}

type clientOpts struct {
	logger    yalogi.Logger
	closeConn bool
}

var defaultClientOpts = clientOpts{
	logger:    yalogi.LogNull,
	closeConn: true,
}

// ClientOption encapsules options for client
type ClientOption func(*clientOpts)

// CloseConnection option closes grpc connection on shutdown
func CloseConnection(b bool) ClientOption {
	return func(o *clientOpts) {
		o.closeConn = b
	}
}

// SetLogger option allows set a custom logger
func SetLogger(l yalogi.Logger) ClientOption {
	return func(o *clientOpts) {
		if l != nil {
			o.logger = l
		}
	}
}

// NewClient returns a new client
func NewClient(conn *grpc.ClientConn, opt ...ClientOption) *Client {
	opts := defaultClientOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Client{
		opts:   opts,
		logger: opts.logger,
		conn:   conn,
		client: pb.NewForwardClient(conn),
	}
}

// ForwardEvent implements event.Forwarder interface
func (c *Client) ForwardEvent(ctx context.Context, e event.Event) error {
	if c.closed {
		return event.ErrUnavailable
	}
	//create request
	req, err := encoding.ForwardEventRequest(e)
	if err != nil {
		return c.mapError(event.ErrBadRequest)
	}
	//notify request
	_, err = c.client.ForwardEvent(ctx, req)
	if err != nil {
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
		return event.ErrCanceledRequest
	case codes.InvalidArgument:
		return event.ErrBadRequest
	case codes.Unimplemented:
		return event.ErrNotSupported
	case codes.Internal:
		return event.ErrInternal
	case codes.Unavailable:
		return event.ErrUnavailable
	default:
		return event.ErrUnavailable
	}
}

//Close the client
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
