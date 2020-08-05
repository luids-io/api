// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	"google.golang.org/grpc"

	"github.com/luids-io/api/netutil/grpc/pb"
)

type pcktClientStream interface {
	grpc.ClientStream
	Send(*pb.SendPacketRequest) error
}

type rpcClient struct {
	client    pb.AnalyzeClient
	stream    pcktClientStream
	dataCh    chan *pb.SendPacketRequest
	connected bool
}

func newRPCClient(c pb.AnalyzeClient, buffSize int) *rpcClient {
	r := &rpcClient{
		client: c,
		dataCh: make(chan *pb.SendPacketRequest, buffSize),
	}
	return r
}

// Send write request in channel
func (r *rpcClient) Send(req *pb.SendPacketRequest) error {
	r.dataCh <- req
	return nil
}

func (r *rpcClient) run(wg *sync.WaitGroup, closeCh <-chan struct{}, errCh chan<- error) {
PROCESSLOOP:
	for {
		select {
		case data := <-r.dataCh:
			err := r.send(data)
			if err != nil {
				errCh <- err
			}
		case <-closeCh:
			//clean buffer
			for data := range r.dataCh {
				err := r.send(data)
				if err != nil {
					errCh <- err
				}
			}
			break PROCESSLOOP
		}
	}
	//close channel data and close stream
	close(r.dataCh)
	r.close()

	wg.Done()
}

//send request, implements a reconnection system
func (r *rpcClient) send(req *pb.SendPacketRequest) error {
	if !r.connected {
		err := r.connect()
		if err != nil {
			return fmt.Errorf("connecting error: %v", err)
		}
	}
	//send
	err := r.stream.Send(req)
	if err != nil {
		if err == io.EOF {
			err = errors.New("connection is closed")
		} else {
			err = fmt.Errorf("sending error: %v", err)
		}
		r.close()
	}
	return err
}

func (r *rpcClient) connect() error {
	var err error
	r.stream, err = r.client.SendPackets(context.Background())
	if err != nil {
		return err
	}
	r.connected = true
	return nil
}

func (r *rpcClient) close() {
	if r.connected {
		r.stream.CloseSend()
		r.connected = false
	}
}
