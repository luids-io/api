// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes"
	"github.com/luids-io/api/netutil"
	"github.com/luids-io/api/netutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Client provides a grpc client.
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

// ClientOption encapsules options for client.
type ClientOption func(*clientOpts)

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

// SetPacketBuffer option sets packet buffer.
func SetPacketBuffer(i int) ClientOption {
	return func(o *clientOpts) {
		if i > 0 {
			o.buffSize = i
		}
	}
}

// NewClient returns a new client.
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

// SendPacket implements netutil.Analyzer interface.
func (c *Client) SendPacket(layer netutil.Layer, md netutil.PacketMetadata, data []byte) error {
	if !c.started {
		c.logger.Warnf("client.netutil.analyze: sendpacket(): client is closed")
		return netutil.ErrUnavailable
	}
	ts, _ := ptypes.TimestampProto(md.Timestamp)
	req := &pb.SendPacketRequest{
		Layer: pb.Layer(layer),
		Data:  data,
		Metadata: &pb.PacketMetadata{
			Timestamp:      ts,
			CaptureLength:  int32(md.CaptureLength),
			Length:         int32(md.Length),
			InterfaceIndex: int32(md.InterfaceIndex),
		},
	}
	err := c.rpc.Send(req)
	if err != nil {
		c.logger.Warnf("client.netutil.analyze: sendpacket(layer=[%v],md=[%v]): %v", layer, md, err)
		return c.mapError(err)
	}
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
	for err := range c.errs {
		c.logger.Warnf("client.netutil.analyze: sendpacket(): %v", err)
	}
}

//Close closes the client.
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

// Ping checks connectivity with the api.
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

//mapping errors
func (c *Client) mapError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.InvalidArgument:
		return netutil.ErrBadRequest
	case codes.OutOfRange:
		if st.Message() == netutil.ErrPacketOutOfOrder.Error() {
			return netutil.ErrPacketOutOfOrder
		}
		return netutil.ErrTimeOutOfSync
	case codes.Unimplemented:
		return netutil.ErrNotSupported
	case codes.ResourceExhausted:
		return netutil.ErrAnalyzerExists
	case codes.Internal:
		return netutil.ErrInternal
	case codes.Unavailable:
		return netutil.ErrUnavailable
	default:
		return netutil.ErrUnavailable
	}
}

//API returns API service name implemented.
func (c *Client) API() string {
	return ServiceName()
}
