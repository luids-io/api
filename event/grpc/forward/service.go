// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package forward

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/event"
	"github.com/luids-io/api/event/grpc/encoding"
	"github.com/luids-io/api/event/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	logger    yalogi.Logger
	forwarder event.Forwarder
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
func NewService(f event.Forwarder, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{forwarder: f, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterForwardServer(server, service)
}

// ForwardEvent implements API service
func (s *Service) ForwardEvent(ctx context.Context, in *pb.ForwardEventRequest) (*empty.Empty, error) {
	e, err := encoding.FromForwardEventRequest(in)
	if err != nil {
		s.logger.Warnf("service.event.forward: [peer=%s] forward(%v,%s): %v", getPeerAddr(ctx), e.Code, e.ID, err)
		return nil, s.mapError(event.ErrBadRequest)
	}
	err = s.forwarder.ForwardEvent(ctx, e)
	if err != nil {
		s.logger.Warnf("service.event.forward: [peer=%s] forward(%v,%s): %v", getPeerAddr(ctx), e.Code, e.ID, err)
		return nil, s.mapError(err)
	}
	return &empty.Empty{}, nil
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
