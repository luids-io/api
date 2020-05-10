// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package check

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/xlist"
	"github.com/luids-io/api/xlist/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service provides a wrapper for the interface xlist.Checker that handles
// grpc requests.
type Service struct {
	opts    serviceOpts
	logger  yalogi.Logger
	checker xlist.Checker
}

type serviceOpts struct {
	logger        yalogi.Logger
	exposePing    bool
	disclosureErr bool
}

var defaultServiceOpts = serviceOpts{
	logger: yalogi.LogNull,
}

// ServiceOption is used for service configuration
type ServiceOption func(*serviceOpts)

// ExposePing exposes ping to the list in the service, allowing not only
// connectivity check
func ExposePing(b bool) ServiceOption {
	return func(o *serviceOpts) {
		o.exposePing = b
	}
}

// DisclosureErrors returns errors without replacing by a generic message
func DisclosureErrors(b bool) ServiceOption {
	return func(o *serviceOpts) {
		o.disclosureErr = b
	}
}

// SetServiceLogger option allows set a custom logger
func SetServiceLogger(l yalogi.Logger) ServiceOption {
	return func(o *serviceOpts) {
		if l != nil {
			o.logger = l
		}
	}
}

// NewService returns a new Service for the cheker
func NewService(checker xlist.Checker, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{checker: checker, opts: opts, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterCheckServer(server, service)
}

// Check implements grpc handler for Check
func (s *Service) Check(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponse, error) {
	req := in.GetRequest()
	if req == nil {
		return nil, s.mapError(xlist.ErrBadRequest)
	}
	resp, err := s.checker.Check(ctx, req.GetName(), xlist.Resource(req.GetResource()))
	if err != nil {
		return nil, s.mapError(err)
	}
	reply := &pb.CheckResponse{
		Response: &pb.Response{
			Result: resp.Result,
			Reason: resp.Reason,
			TTL:    int32(resp.TTL),
		}}
	return reply, nil
}

// Resources implements grpc handler for Resources
func (s *Service) Resources(ctx context.Context, in *empty.Empty) (*pb.ResourcesResponse, error) {
	resources := s.checker.Resources()
	retres := make([]pb.Resource, 0, len(resources))
	for _, r := range resources {
		retres = append(retres, pb.Resource(r))
	}
	return &pb.ResourcesResponse{Resources: retres}, nil
}

// Ping implements grpc handler for Ping
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
		rpcerr := status.Error(codes.Internal, xlist.ErrInternal.Error())
		if s.opts.disclosureErr {
			rpcerr = status.Error(codes.Internal, err.Error())
		}
		return rpcerr
	}
}