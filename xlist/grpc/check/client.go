// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package check

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/xlist"
	"github.com/luids-io/api/xlist/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Client provides a grpc client.
type Client struct {
	opts   clientOpts
	logger yalogi.Logger
	//grpc connection
	conn   *grpc.ClientConn
	client pb.CheckClient
	//control
	closed bool
	//cache
	cache *cache
}

// ClientOption encapsules options for client.
type ClientOption func(*clientOpts)

type clientOpts struct {
	logger    yalogi.Logger
	closeConn bool
	//cache opts
	useCache     bool
	minttl       int
	maxttl       int
	cacheCleanup time.Duration
}

var defaultClientOpts = clientOpts{
	logger:       yalogi.LogNull,
	closeConn:    true,
	cacheCleanup: defaultCacheCleanups,
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

// SetCache sets cache ttl and negative ttl.
func SetCache(minttl, maxttl int) ClientOption {
	return func(o *clientOpts) {
		if minttl >= 0 && maxttl >= 0 {
			o.minttl = minttl
			o.maxttl = maxttl
			o.useCache = true
		}
	}
}

// SetCacheCleanUps sets interval between cache cleanups.
func SetCacheCleanUps(d time.Duration) ClientOption {
	return func(o *clientOpts) {
		if d > 0 {
			o.cacheCleanup = d
		}
	}
}

// NewClient returns a new Client.
func NewClient(conn *grpc.ClientConn, opt ...ClientOption) *Client {
	opts := defaultClientOpts
	for _, o := range opt {
		o(&opts)
	}
	c := &Client{
		opts:   opts,
		logger: opts.logger,
		conn:   conn,
		client: pb.NewCheckClient(conn),
	}
	if opts.useCache {
		c.cache = newCache(opts.minttl, opts.maxttl, opts.cacheCleanup)
	}
	return c
}

// Check implements xlist.Checker interface.
func (c *Client) Check(ctx context.Context, name string, resource xlist.Resource) (xlist.Response, error) {
	if c.closed {
		c.logger.Warnf("client.xlist.check: check(%s,%v): client is closed", name, resource)
		return xlist.Response{}, xlist.ErrUnavailable
	}
	name, ctx, err := xlist.DoValidation(ctx, name, resource, false)
	if err != nil {
		c.logger.Warnf("client.xlist.check: check(%s,%v): %v", name, resource, err)
		return xlist.Response{}, err
	}
	if c.opts.useCache {
		resp, ok := c.cache.get(name, resource)
		if ok {
			return resp, nil
		}
	}
	resp, err := c.doCheck(ctx, name, resource)
	if c.opts.useCache && err == nil {
		c.cache.set(name, resource, resp)
	}
	return resp, err
}

// Resources implements xlist.Checker interface.
func (c *Client) Resources(ctx context.Context) ([]xlist.Resource, error) {
	if c.closed {
		c.logger.Warnf("client.xlist.check: resources(): client is closed")
		return nil, xlist.ErrUnavailable
	}
	resp, err := c.client.Resources(ctx, &empty.Empty{})
	if err != nil {
		c.logger.Warnf("client.xlist.check: resources(): %v", err)
		return []xlist.Resource{}, c.mapError(err)
	}
	resources := make([]xlist.Resource, 0, len(resp.Resources))
	for _, r := range resp.Resources {
		resources = append(resources, xlist.Resource(r))
	}
	return resources, nil
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

func (c *Client) doCheck(ctx context.Context, name string, resource xlist.Resource) (xlist.Response, error) {
	req := &pb.CheckRequest{Name: name, Resource: pb.Resource(resource)}
	res, err := c.client.Check(ctx, req)
	if err != nil {
		c.logger.Warnf("client.xlist.check: check(%s,%v): %v", name, resource, err)
		return xlist.Response{}, c.mapError(err)
	}
	r := xlist.Response{
		Result: res.GetResult(),
		Reason: res.GetReason(),
		TTL:    int(res.GetTTL())}
	return r, nil
}

//mapping errors
func (c *Client) mapError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.Canceled:
		return xlist.ErrCanceledRequest
	case codes.InvalidArgument:
		return xlist.ErrBadRequest
	case codes.Unimplemented:
		return xlist.ErrNotSupported
	case codes.Internal:
		return xlist.ErrInternal
	case codes.Unavailable:
		return xlist.ErrUnavailable
	default:
		return xlist.ErrUnavailable
	}
}

// Flush cache if set.
func (c *Client) Flush() {
	if !c.closed && c.opts.useCache {
		c.cache.flush()
	}
}

// Close the client.
func (c *Client) Close() error {
	if c.closed {
		return errors.New("client closed")
	}
	c.closed = true
	if c.opts.useCache {
		c.cache.flush()
		c.cache = nil
	}
	if c.opts.closeConn {
		return c.conn.Close()
	}
	return nil
}

// API returns API service name implemented.
func (c *Client) API() string {
	return ServiceName()
}
