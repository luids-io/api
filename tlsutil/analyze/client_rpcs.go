// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/luids-io/api/protogen/tlsutilpb"
)

type pcktClientStream interface {
	grpc.ClientStream
	Send(*pb.SendMessageRequest) error
}

type rpcClient struct {
	client    pb.AnalyzeClient
	stream    pcktClientStream
	dataCh    chan *pb.SendMessageRequest
	connected bool
}

func newRPCClient(c pb.AnalyzeClient, buffSize int) *rpcClient {
	r := &rpcClient{
		client: c,
		dataCh: make(chan *pb.SendMessageRequest, buffSize),
	}
	return r
}

// Data returns channel for write data
func (r *rpcClient) Data() chan<- *pb.SendMessageRequest {
	return r.dataCh
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
func (r *rpcClient) send(req *pb.SendMessageRequest) error {
	if !r.connected {
		err := r.connect()
		if err != nil {
			return fmt.Errorf("connecting error: %v", err)
		}
	}
	//send
	err := r.stream.Send(req)
	if err != nil {
		err = fmt.Errorf("sending error: %v", err)
		r.close()
	}
	return err
}

func (r *rpcClient) connect() error {
	var err error
	r.stream, err = r.client.SendMessages(context.Background())
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
