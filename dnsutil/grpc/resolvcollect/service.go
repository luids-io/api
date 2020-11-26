// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package resolvcollect

import (
	"context"
	"errors"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a grpc service wrapper.
type Service struct {
	logger    yalogi.Logger
	collector dnsutil.ResolvCollector
}

// ServiceOption is used for service configuration
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
func NewService(c dnsutil.ResolvCollector, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{collector: c, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server.
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterResolvCollectServer(server, service)
}

// Collect implements grpc api.
func (s *Service) Collect(ctx context.Context, req *pb.ResolvCollectRequest) (*empty.Empty, error) {
	//parse request
	client, name, resolved, cnames, err := parseRequest(req)
	if err != nil {
		s.logger.Warnf("service.dnsutil.resolvcollect: [peer=%s] collect(%v,%s,%v,%v): %v", getPeerAddr(ctx), client, name, resolved, cnames, err)
		return nil, s.mapError(dnsutil.ErrBadRequest)
	}
	//do request
	err = s.collector.Collect(ctx, client, name, resolved, cnames)
	if err != nil {
		s.logger.Warnf("service.dnsutil.resolvcollect: [peer=%s] collect(%v,%s,%v,%v): %v", getPeerAddr(ctx), client, name, resolved, cnames, err)
		return nil, s.mapError(err)
	}
	//return response
	return &empty.Empty{}, nil
}

func parseRequest(req *pb.ResolvCollectRequest) (net.IP, string, []net.IP, []string, error) {
	client := net.ParseIP(req.GetClientIp())
	if client == nil {
		return nil, "", nil, nil, errors.New("bad client ip")
	}
	name := req.GetName()
	if name == "" {
		return nil, "", nil, nil, errors.New("bad dns name")
	}
	if len(req.GetResolvedIps()) == 0 {
		return nil, "", nil, nil, errors.New("resolved ips empty")
	}
	resolved := make([]net.IP, 0, len(req.GetResolvedIps()))
	for _, r := range req.GetResolvedIps() {
		ip := net.ParseIP(r)
		if ip == nil {
			return nil, "", nil, nil, errors.New("bad resolved ip")
		}
		resolved = append(resolved, ip)
	}
	cnames := make([]string, 0, len(req.GetResolvedCnames()))
	for _, r := range req.GetResolvedCnames() {
		cnames = append(cnames, r)
	}
	return client, name, resolved, cnames, nil
}

//mapping checking errors
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
	case dnsutil.ErrLimitDNSClientQueries:
		return status.Error(codes.ResourceExhausted, err.Error())
	case dnsutil.ErrLimitResolvedNamesIP:
		return status.Error(codes.ResourceExhausted, err.Error())
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
