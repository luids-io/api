// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package archive

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
	archiver event.Archiver
}

// NewService returns a new Service for the grpc api
func NewService(a event.Archiver) *Service {
	return &Service{archiver: a}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterArchiveServer(server, service)
}

// SaveEvent implements API service
func (s *Service) SaveEvent(ctx context.Context, in *pb.SaveEventRequest) (*pb.SaveEventResponse, error) {
	e, err := encoding.FromSaveEventRequest(in)
	if err != nil {
		rpcerr := status.Error(codes.InvalidArgument, "request is not valid")
		return nil, rpcerr
	}
	sID, err := s.archiver.SaveEvent(ctx, e)
	if err != nil {
		rpcerr := status.Error(codes.Internal, err.Error())
		return nil, rpcerr
	}
	reply := &pb.SaveEventResponse{StorageID: sID}
	return reply, nil
}

//mapping errors
func (s *Service) mapError(err error) error {
	return status.Error(codes.Unavailable, err.Error())
}
