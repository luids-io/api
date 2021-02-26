// Copyright 2021 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package finder

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
	logger yalogi.Logger
	finder dnsutil.Finder
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
func NewService(f dnsutil.Finder, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{finder: f, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server.
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterFinderServer(server, service)
}

// GetResolv implements grpc interface.
func (s *Service) GetResolv(ctx context.Context, req *pb.GetResolvRequest) (*pb.GetResolvResponse, error) {
	id := req.GetId()
	if id == "" {
		s.logger.Warnf("service.dnsutil.finder: [peer=%s] getresolv(): id is empty", getPeerAddr(ctx))
		return nil, s.mapError(dnsutil.ErrBadRequest)
	}
	data, exists, err := s.finder.GetResolv(ctx, id)
	if err != nil {
		s.logger.Warnf("service.dnsutil.finder: [peer=%s] getresolv(%s): %v", getPeerAddr(ctx), id, err)
		return nil, s.mapError(err)
	}
	if !exists {
		return nil, status.Error(codes.NotFound, "resolv not found")
	}
	datapb, err := encoding.ResolvDataPB(&data)
	if err != nil {
		s.logger.Warnf("service.dnsutil.finder: [peer=%s] getresolv(%s): encoding pb: %v", getPeerAddr(ctx), id, err)
		return nil, s.mapError(err)
	}
	return &pb.GetResolvResponse{Data: datapb}, nil
}

// ListResolvs implements grpc interface.
func (s *Service) ListResolvs(ctx context.Context, req *pb.ListResolvsRequest) (*pb.ListResolvsResponse, error) {
	// get request
	max := int(req.GetMax())
	if max < 0 {
		s.logger.Warnf("service.dnsutil.finder: [peer=%s] listresolvs(): bad max", getPeerAddr(ctx))
		return nil, s.mapError(dnsutil.ErrBadRequest)
	}
	reverse := req.GetReverse()
	filters := make([]dnsutil.ResolvsFilter, 0, len(req.GetFilters()))
	for _, f := range req.GetFilters() {
		nf, err := encoding.ResolvsFilter(f)
		if err != nil {
			s.logger.Warnf("service.dnsutil.finder: [peer=%s] listresolvs(): bad filter: %v", getPeerAddr(ctx), err)
			return nil, s.mapError(dnsutil.ErrBadRequest)
		}
		filters = append(filters, nf)
	}
	//do list
	data, next, err := s.finder.ListResolvs(ctx, filters, reverse, max, req.GetNext())
	if err != nil {
		s.logger.Warnf("service.dnsutil.finder: [peer=%s] listresolvs(): %v", getPeerAddr(ctx), err)
		return nil, s.mapError(err)
	}
	//prepare response
	resp := &pb.ListResolvsResponse{
		Next: next,
		Data: make([]*pb.ResolvData, 0, len(data)),
	}
	for _, r := range data {
		pbdata, err := encoding.ResolvDataPB(&r)
		if err != nil {
			s.logger.Warnf("service.dnsutil.finder: [peer=%s] listresolvs(): %v", getPeerAddr(ctx), err)
			return nil, s.mapError(err)
		}
		resp.Data = append(resp.Data, pbdata)
	}
	return resp, nil
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
