// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package resolvcheck

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
	"github.com/luids-io/api/xlist"
	"github.com/luids-io/core/yalogi"
)

// Client provides a client for resolvcheck api
type Client struct {
	opts   clientOpts
	logger yalogi.Logger
	//grpc connection
	conn   *grpc.ClientConn
	client pb.ResolvCheckClient
	//control
	closed bool
	//cache
	cache *cache
}

type clientOpts struct {
	logger    yalogi.Logger
	closeConn bool
	//client mapping
	cmap *ClientMap
	//cache opts
	useCache     bool
	ttl          int
	negativettl  int
	cacheCleanup time.Duration
}

var defaultClientOpts = clientOpts{
	logger:       yalogi.LogNull,
	closeConn:    true,
	cmap:         NewClientMap(),
	cacheCleanup: defaultCacheCleanups,
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

// SetClientMap option allows set a client mapper
func SetClientMap(cmap *ClientMap) ClientOption {
	return func(o *clientOpts) {
		if cmap != nil {
			o.cmap = cmap
		}
	}
}

// SetCache sets cache ttl and negative ttl
func SetCache(ttl, negativettl int) ClientOption {
	return func(o *clientOpts) {
		if ttl >= xlist.NeverCache && negativettl >= xlist.NeverCache {
			o.ttl = ttl
			o.negativettl = negativettl
			o.useCache = true
		}
	}
}

// SetCacheCleanUps sets interval between cache cleanups
func SetCacheCleanUps(d time.Duration) ClientOption {
	return func(o *clientOpts) {
		if d > 0 {
			o.cacheCleanup = d
		}
	}
}

// NewClient returns a new Client
func NewClient(conn *grpc.ClientConn, opt ...ClientOption) *Client {
	opts := defaultClientOpts
	for _, o := range opt {
		o(&opts)
	}
	c := &Client{
		opts:   opts,
		logger: opts.logger,
		conn:   conn,
		client: pb.NewResolvCheckClient(conn),
	}
	if opts.useCache {
		c.cache = newCache(opts.ttl, opts.negativettl, opts.cacheCleanup)
	}
	return c
}

// Check implements dnsutil.ResolvChecker interface
func (c *Client) Check(ctx context.Context, client, resolved net.IP, name string) (dnsutil.ResolvResponse, error) {
	if c.closed {
		return dnsutil.ResolvResponse{}, dnsutil.ErrBadRequest
	}
	if c.opts.useCache {
		resp, ok := c.cache.get(client, resolved, name)
		if ok {
			return resp, nil
		}
	}
	//exec query
	response, err := c.doCheck(ctx, client, resolved, name)
	if c.opts.useCache {
		if err == nil {
			c.cache.set(client, resolved, name, response)
		}
	}
	return response, err
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

func (c *Client) doCheck(ctx context.Context, client, resolved net.IP, name string) (dnsutil.ResolvResponse, error) {
	//exec query
	response, err := c.client.Check(ctx,
		&pb.ResolvCheckRequest{
			ClientIp:   c.opts.cmap.Get(client).String(),
			ResolvedIp: resolved.String(),
			Name:       name,
		})
	if err != nil {
		return dnsutil.ResolvResponse{}, c.mapError(err)
	}
	//process response
	resp := dnsutil.ResolvResponse{}
	resp.Result = response.GetResult()
	tstamp := response.GetLastTs()
	if tstamp != nil {
		resp.Last, _ = ptypes.Timestamp(tstamp)
	}
	tstamp = response.GetStoreTs()
	if tstamp != nil {
		resp.Store, _ = ptypes.Timestamp(tstamp)
	}
	return resp, nil
}

//API returns API service name implemented
func (c *Client) API() string {
	return ServiceName()
}
