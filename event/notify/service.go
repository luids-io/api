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
	"github.com/luids-io/core/utils/yalogi"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	logger   yalogi.Logger
	notifier event.Notifier
}

type serviceOpts struct {
	logger yalogi.Logger
}

var defaultServiceOpts = serviceOpts{logger: yalogi.LogNull}

// ServiceOption is used for service configuration
type ServiceOption func(*serviceOpts)

// SetServiceLogger option allows set a custom logger
func SetServiceLogger(l yalogi.Logger) ServiceOption {
	return func(o *serviceOpts) {
		if l != nil {
			o.logger = l
		}
	}
}

// NewService returns a new Service for the grpc api
func NewService(n event.Notifier, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{notifier: n, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterNotifyServer(server, service)
}

// NotifyEvent implements API service
func (s *Service) NotifyEvent(ctx context.Context, in *pb.NotifyEventRequest) (*pb.NotifyEventResponse, error) {
	e, err := encoding.FromNotifyEventRequest(in)
	if err != nil {
		return nil, s.mapError(event.ErrBadRequest)
	}
	eventID, err := s.notifier.NotifyEvent(ctx, e)
	if err != nil {
		return nil, s.mapError(err)
	}
	return &pb.NotifyEventResponse{EventID: eventID}, nil
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
