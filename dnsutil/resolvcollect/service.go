// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package resolvcollect

import (
	"context"
	"errors"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/luids-io/api/protogen/dnsutilpb"
	"github.com/luids-io/core/dnsutil"
	"github.com/luids-io/core/utils/yalogi"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	logger    yalogi.Logger
	collector dnsutil.ResolvCollector
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
func NewService(c dnsutil.ResolvCollector, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{collector: c, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterResolvCollectServer(server, service)
}

// Collect implements interface
func (s *Service) Collect(ctx context.Context, req *pb.ResolvCollectRequest) (*empty.Empty, error) {
	//parse request
	client, name, resolved, err := parseRequest(req)
	if err != nil {
		return nil, s.mapError(dnsutil.ErrBadRequest)
	}
	//do request
	err = s.collector.Collect(ctx, client, name, resolved)
	if err != nil {
		return nil, s.mapError(err)
	}
	//return response
	return &empty.Empty{}, nil
}

func parseRequest(req *pb.ResolvCollectRequest) (net.IP, string, []net.IP, error) {
	client := net.ParseIP(req.GetClientIp())
	if client == nil {
		return nil, "", nil, errors.New("bad client ip")
	}
	name := req.GetName()
	if name == "" {
		return nil, "", nil, errors.New("bad dns name")
	}
	if len(req.GetResolvedIps()) == 0 {
		return nil, "", nil, errors.New("resolved ips empty")
	}
	resolved := make([]net.IP, 0, len(req.GetResolvedIps()))
	for _, r := range req.GetResolvedIps() {
		ip := net.ParseIP(r)
		if ip == nil {
			return nil, "", nil, errors.New("bad resolved ip")
		}
		resolved = append(resolved, ip)
	}
	return client, name, resolved, nil
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
