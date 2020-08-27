// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package classify

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/api/tlsutil/grpc/encoding"
	"github.com/luids-io/api/tlsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Client provides a grpc client.
type Client struct {
	opts   clientOpts
	logger yalogi.Logger
	//grpc connection
	conn   *grpc.ClientConn
	client pb.ClassifyClient
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

// CloseConnection option closes grpc connection on close.
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

// NewClient returns a new grpc Client.
func NewClient(conn *grpc.ClientConn, opt ...ClientOption) *Client {
	opts := defaultClientOpts
	for _, o := range opt {
		o(&opts)
	}
	c := &Client{
		opts:   opts,
		logger: opts.logger,
		conn:   conn,
		client: pb.NewClassifyClient(conn),
	}
	return c
}

// ClassifyConnections implements tlsutil.Classifier.
func (c *Client) ClassifyConnections(ctx context.Context, requests []*tlsutil.ConnectionData) ([]tlsutil.ClassifyResponse, error) {
	if c.closed {
		c.logger.Warnf("client.tlsutil.classify: connections(#%v): client is closed", len(requests))
		return nil, tlsutil.ErrUnavailable
	}
	if len(requests) == 0 {
		c.logger.Warnf("client.tlsutil.classify: connections(): requests len can't be empty")
		return nil, tlsutil.ErrBadRequest
	}
	// prepare requests
	sendRequests := make([]*pb.ConnectionData, 0, len(requests))
	for _, cdata := range requests {
		sendRequests = append(sendRequests, encoding.ConnectionDataPB(cdata))
	}
	// do classify
	pbres, err := c.client.Connections(ctx, &pb.ClassifyConnectionsRequest{Connections: sendRequests})
	if err != nil {
		c.logger.Warnf("client.tlsutil.classify: connections(#%v): %v", len(requests), err)
		return nil, c.mapError(err)
	}
	if len(requests) != len(pbres.Responses) {
		c.logger.Warnf("client.tlsutil.classify: connections(#%v): "+
			"requests len and responses len (#%v) missmatch", len(requests), len(pbres.Responses))
		return nil, tlsutil.ErrInternal
	}
	// reencode responses
	responses := make([]tlsutil.ClassifyResponse, 0, len(pbres.Responses))
	for _, r := range pbres.Responses {
		resp := tlsutil.ClassifyResponse{}
		//resp.ID
		if r.GetErr() != "" {
			resp.Err = errors.New(r.GetErr())
		} else {
			resp.Results = make([]tlsutil.ClassifyResult, 0, len(r.GetResults()))
			for _, result := range r.GetResults() {
				resp.Results = append(resp.Results, tlsutil.ClassifyResult{
					Label: result.GetLabel(),
					Prob:  result.GetProb(),
				})
			}
		}
		responses = append(responses, resp)
	}
	return responses, nil
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

//mapping errors
func (c *Client) mapError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.InvalidArgument:
		return tlsutil.ErrBadRequest
	case codes.Unimplemented:
		return tlsutil.ErrNotSupported
	case codes.Internal:
		return tlsutil.ErrInternal
	case codes.Unavailable:
		return tlsutil.ErrUnavailable
	default:
		return tlsutil.ErrUnavailable
	}
}
