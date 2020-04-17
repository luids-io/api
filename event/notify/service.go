// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package notify

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/event/encoding"
	pb "github.com/luids-io/api/protogen/eventpb"
	"github.com/luids-io/core/event"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	notifier event.Notifier
}

// NewService returns a new Service for the grpc api
func NewService(n event.Notifier) *Service {
	return &Service{notifier: n}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterNotifyServer(server, service)
}

// NotifyEvent implements API service
func (s *Service) NotifyEvent(ctx context.Context, in *pb.NotifyEventRequest) (*pb.NotifyEventResponse, error) {
	e, err := encoding.FromNotifyEventRequest(in)
	if err != nil {
		rpcerr := status.Error(codes.InvalidArgument, "request is not valid")
		return nil, rpcerr
	}
	eventID, err := s.notifier.NotifyEvent(ctx, e)
	if err != nil {
		rpcerr := status.Error(codes.Internal, err.Error())
		return nil, rpcerr
	}
	return &pb.NotifyEventResponse{EventID: eventID}, nil
}

//mapping errors
func (s *Service) mapError(err error) error {
	return status.Error(codes.Unavailable, err.Error())
}
