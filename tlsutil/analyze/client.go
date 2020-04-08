// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	pb "github.com/luids-io/api/protogen/tlsutilpb"
	"github.com/luids-io/api/tlsutil/encoding"
	"github.com/luids-io/core/tlsutil"
	"github.com/luids-io/core/utils/yalogi"
)

// Client is the main struct for grpc client
type Client struct {
	opts   clientOpts
	logger yalogi.Logger
	//grpc connection
	conn   *grpc.ClientConn
	client pb.AnalyzeClient
	//rpc management
	rpc *rpcClient
	//control
	started bool
	close   chan struct{}
	errs    chan error
	wg      sync.WaitGroup
}

type clientOpts struct {
	logger    yalogi.Logger
	closeConn bool
	buffSize  int
}

var defaultClientOpts = clientOpts{
	logger:    yalogi.LogNull,
	closeConn: true,
	buffSize:  100,
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
	c := &Client{
		opts:   opts,
		logger: opts.logger,
		conn:   conn,
		client: pb.NewAnalyzeClient(conn),
	}
	c.start()
	return c
}

// SendMessage implements capture.Analyzer interface
func (c *Client) SendMessage(msg *tlsutil.Msg) error {
	if !c.started {
		return errors.New("client not started")
	}
	c.rpc.Data() <- encoding.MessageRequestPB(msg)
	return nil
}

func (c *Client) start() {
	//init status
	c.close = make(chan struct{})
	c.errs = make(chan error, c.opts.buffSize)
	go c.processErrs()

	//init rpc manager
	c.wg.Add(1)
	c.rpc = newRPCClient(c.client, c.opts.buffSize)
	go c.rpc.run(&c.wg, c.close, c.errs)

	c.started = true
}

func (c *Client) processErrs() {
	for e := range c.errs {
		c.logger.Warnf("%v", e)
	}
}

//Close closes the client
func (c *Client) Close() error {
	if !c.started {
		return errors.New("client closed")
	}
	c.started = false
	//close signal
	close(c.close)
	c.wg.Wait()
	close(c.errs)

	if c.opts.closeConn {
		return c.conn.Close()
	}
	return nil
}

// Ping checks connectivity with the api
func (c *Client) Ping() error {
	if !c.started {
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
