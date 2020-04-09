// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	pb "github.com/luids-io/api/protogen/tlsutilpb"
	"github.com/luids-io/api/tlsutil/encoding"
	"github.com/luids-io/core/tlsutil"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	factory tlsutil.AnalyzerFactory
}

// NewService returns a new Service for the grpc api
func NewService(f tlsutil.AnalyzerFactory) *Service {
	return &Service{factory: f}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterAnalyzeServer(server, service)
}

// SendMessages manage messages
func (s *Service) SendMessages(stream pb.Analyze_SendMessagesServer) error {
	ctx := stream.Context()
	p, ok := peer.FromContext(ctx)
	if !ok {
		return status.Errorf(codes.Internal, "Internal error getting peer")
	}
	// creates packet source
	name := p.Addr.String()
	analyzer, err := s.factory.NewAnalyzer(name)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error getting analyzer")
	}
	defer analyzer.Close()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		msg := encoding.MessageRequest(req)
		analyzer.SendMessage(msg)
	}

}

//mapping errors
func (s *Service) mapError(err error) error {
	return status.Error(codes.Unavailable, err.Error())
}
