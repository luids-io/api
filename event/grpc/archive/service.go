// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package archive

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/event"
	"github.com/luids-io/api/event/grpc/encoding"
	"github.com/luids-io/api/event/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a grpc service wrapper.
type Service struct {
	logger   yalogi.Logger
	archiver event.Archiver
}

// ServiceOption is used for service configuration.
type ServiceOption func(*serviceOpts)

type serviceOpts struct {
	logger yalogi.Logger
}

var defaultServiceOpts = serviceOpts{logger: yalogi.LogNull}

// SetServiceLogger option allows set a custom logger.
func SetServiceLogger(l yalogi.Logger) ServiceOption {
	return func(o *serviceOpts) {
		if l != nil {
			o.logger = l
		}
	}
}

// NewService returns a new Service.
func NewService(a event.Archiver, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{archiver: a, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server.
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterArchiveServer(server, service)
}

// SaveEvent implements API service.
func (s *Service) SaveEvent(ctx context.Context, in *pb.SaveEventRequest) (*pb.SaveEventResponse, error) {
	e, err := encoding.FromSaveEventRequest(in)
	if err != nil {
		s.logger.Warnf("service.event.archive: [peer=%s] save(%s): %v", getPeerAddr(ctx), e.ID, err)
		return nil, s.mapError(event.ErrBadRequest)
	}
	sID, err := s.archiver.SaveEvent(ctx, e)
	if err != nil {
		s.logger.Warnf("service.event.archive: [peer=%s] save(%s): %v", getPeerAddr(ctx), e.ID, err)
		return nil, s.mapError(err)
	}
	return &pb.SaveEventResponse{StorageID: sID}, nil
}

//mapping errors
func (s *Service) mapError(err error) error {
	switch err {
	case event.ErrCanceledRequest:
		return status.Error(codes.Canceled, err.Error())
	case event.ErrBadRequest:
		return status.Error(codes.InvalidArgument, err.Error())
	case event.ErrUnauthorized:
		return status.Error(codes.PermissionDenied, err.Error())
	case event.ErrNotSupported:
		return status.Error(codes.Unimplemented, err.Error())
	case event.ErrUnavailable:
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Internal, event.ErrInternal.Error())
	}
}

func getPeerAddr(ctx context.Context) (paddr string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		paddr = p.Addr.String()
	}
	return
}
