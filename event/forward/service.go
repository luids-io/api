// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package forward

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/luids-io/api/event/encoding"
	pb "github.com/luids-io/api/protogen/eventpb"
	"github.com/luids-io/core/event"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	forwarder event.Forwarder
}

// NewService returns a new Service for the grpc api
func NewService(f event.Forwarder) *Service {
	return &Service{forwarder: f}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterForwardServer(server, service)
}

// ForwardEvent implements API service
func (s *Service) ForwardEvent(ctx context.Context, in *pb.ForwardEventRequest) (*empty.Empty, error) {
	e, err := encoding.FromForwardEventRequest(in)
	if err != nil {
		return &empty.Empty{}, status.Error(codes.InvalidArgument, "request is not valid")
	}
	err = s.forwarder.ForwardEvent(ctx, e)
	if err != nil {
		return &empty.Empty{}, status.Error(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

//mapping errors
func (s *Service) mapError(err error) error {
	return status.Error(codes.Unavailable, err.Error())
}
