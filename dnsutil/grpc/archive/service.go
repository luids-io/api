// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package archive

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/encoding"
	"github.com/luids-io/api/dnsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a grpc service wrapper.
type Service struct {
	logger   yalogi.Logger
	archiver dnsutil.Archiver
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
func NewService(a dnsutil.Archiver, opt ...ServiceOption) *Service {
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

// SaveResolv implements grpc interface.
func (s *Service) SaveResolv(ctx context.Context, req *pb.SaveResolvRequest) (*pb.SaveResolvResponse, error) {
	//get request
	data, err := encoding.FromSaveResolvRequest(req)
	if err != nil {
		s.logger.Warnf("service.dnsutil.archive: [peer=%s] saveresolv(%v,%v): %v", getPeerAddr(ctx), data.Client, data.QID, err)
		return nil, s.mapError(dnsutil.ErrBadRequest)
	}
	//do save
	newid, err := s.archiver.SaveResolv(ctx, data)
	if err != nil {
		s.logger.Warnf("service.dnsutil.archive: [peer=%s] saveresolv(%v,%v): %v", getPeerAddr(ctx), data.Client, data.QID, err)
		return nil, s.mapError(err)
	}
	//return response
	return &pb.SaveResolvResponse{Id: newid}, nil
}

//mapping errors
func (s *Service) mapError(err error) error {
	switch err {
	case dnsutil.ErrCanceledRequest:
		return status.Error(codes.Canceled, err.Error())
	case dnsutil.ErrBadRequest:
		return status.Error(codes.InvalidArgument, err.Error())
	case dnsutil.ErrNotSupported:
		return status.Error(codes.Unimplemented, err.Error())
	case dnsutil.ErrUnavailable:
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Internal, dnsutil.ErrInternal.Error())
	}
}

func getPeerAddr(ctx context.Context) (paddr string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		paddr = p.Addr.String()
	}
	return
}
