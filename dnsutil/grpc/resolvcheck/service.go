// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package resolvcheck

import (
	"context"
	"errors"
	"net"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service provides a wrapper
type Service struct {
	logger  yalogi.Logger
	checker dnsutil.ResolvChecker
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

// NewService returns a new Service
func NewService(c dnsutil.ResolvChecker, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{checker: c, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterResolvCheckServer(server, service)
}

// Check implements API
func (s *Service) Check(ctx context.Context, in *pb.ResolvCheckRequest) (*pb.ResolvCheckResponse, error) {
	//parse request
	client, resolved, name, err := parseRequest(in)
	if err != nil {
		s.logger.Warnf("service.dnsutil.resolvcheck: [peer=%s] check(%v,%v,%s): %v", getPeerAddr(ctx), client, resolved, name, err)
		return nil, s.mapError(dnsutil.ErrBadRequest)
	}
	//do request
	resp, err := s.checker.Check(ctx, client, resolved, name)
	if err != nil {
		s.logger.Warnf("service.dnsutil.resolvcheck: [peer=%s] check(%v,%v,%s): %v", getPeerAddr(ctx), client, resolved, name, err)
		return nil, s.mapError(err)
	}
	//return response
	response := &pb.ResolvCheckResponse{}
	response.Result = resp.Result
	response.LastTs, _ = ptypes.TimestampProto(resp.Last)
	response.StoreTs, _ = ptypes.TimestampProto(resp.Store)
	return response, nil
}

func parseRequest(req *pb.ResolvCheckRequest) (net.IP, net.IP, string, error) {
	client := req.GetClientIp()
	resolved := req.GetResolvedIp()
	if client == "" || resolved == "" {
		return nil, nil, "", errors.New("client and resolved are required")
	}
	clientIP := net.ParseIP(client)
	if clientIP == nil {
		return nil, nil, "", errors.New("client must be an ip")
	}
	resolvedIP := net.ParseIP(resolved)
	if resolvedIP == nil {
		return nil, nil, "", errors.New("resolved must be an ip")
	}
	name := req.GetName()
	return clientIP, resolvedIP, name, nil
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
