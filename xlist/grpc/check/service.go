// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package check

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/xlist"
	"github.com/luids-io/api/xlist/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a grpc service wrapper.
type Service struct {
	opts    serviceOpts
	logger  yalogi.Logger
	checker xlist.Checker
}

// ServiceOption is used for service configuration.
type ServiceOption func(*serviceOpts)

type serviceOpts struct {
	logger        yalogi.Logger
	exposePing    bool
	disclosureErr bool
}

var defaultServiceOpts = serviceOpts{logger: yalogi.LogNull}

// ExposePing exposes ping to the list in the service, allowing not only
// connectivity check.
func ExposePing(b bool) ServiceOption {
	return func(o *serviceOpts) {
		o.exposePing = b
	}
}

// DisclosureErrors returns errors without replacing by a generic message.
func DisclosureErrors(b bool) ServiceOption {
	return func(o *serviceOpts) {
		o.disclosureErr = b
	}
}

// SetServiceLogger option allows set a custom logger.
func SetServiceLogger(l yalogi.Logger) ServiceOption {
	return func(o *serviceOpts) {
		if l != nil {
			o.logger = l
		}
	}
}

// NewService returns a new Service.
func NewService(checker xlist.Checker, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{checker: checker, opts: opts, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server.
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterCheckServer(server, service)
}

// Check implements grpc api.
func (s *Service) Check(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponse, error) {
	//parse request
	name := in.GetName()
	resource := xlist.Resource(in.GetResource())
	//do request
	resp, err := s.checker.Check(ctx, name, resource)
	if err != nil {
		s.logger.Warnf("service.xlist.check: [peer=%s] check(%s,%v): %v", getPeerAddr(ctx), name, resource, err)
		return nil, s.mapError(err)
	}
	//return response
	reply := &pb.CheckResponse{
		Result: resp.Result,
		Reason: resp.Reason,
		TTL:    int32(resp.TTL),
	}
	return reply, nil
}

// Resources implements grpc api.
func (s *Service) Resources(ctx context.Context, in *empty.Empty) (*pb.ResourcesResponse, error) {
	resources := s.checker.Resources()
	retres := make([]pb.Resource, 0, len(resources))
	for _, r := range resources {
		retres = append(retres, pb.Resource(r))
	}
	return &pb.ResourcesResponse{Resources: retres}, nil
}

// Ping implements grpc handler for Ping.
func (s *Service) Ping(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	if s.opts.exposePing {
		err := s.checker.Ping()
		if err != nil {
			rpcerr := status.Error(codes.Unavailable, xlist.ErrUnavailable.Error())
			if s.opts.disclosureErr {
				rpcerr = status.Error(codes.Unavailable, err.Error())
			}
			return nil, rpcerr
		}
	}
	return &empty.Empty{}, nil
}

//mapping checking errors
func (s *Service) mapError(err error) error {
	switch err {
	case xlist.ErrCanceledRequest:
		return status.Error(codes.Canceled, err.Error())
	case xlist.ErrBadRequest:
		return status.Error(codes.InvalidArgument, err.Error())
	case xlist.ErrNotSupported:
		return status.Error(codes.Unimplemented, err.Error())
	case xlist.ErrUnavailable:
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Internal, xlist.ErrInternal.Error())
	}
}

func getPeerAddr(ctx context.Context) (paddr string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		paddr = p.Addr.String()
	}
	return
}
