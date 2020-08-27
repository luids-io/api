// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package classify

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/api/tlsutil/grpc/encoding"
	"github.com/luids-io/api/tlsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a grpc service wrapper.
type Service struct {
	logger     yalogi.Logger
	classifier tlsutil.Classifier
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

// NewService returns a new Service for the cheker.
func NewService(c tlsutil.Classifier, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{classifier: c, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server.
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterClassifyServer(server, service)
}

// Connections implements grpc api.
func (s *Service) Connections(ctx context.Context, in *pb.ClassifyConnectionsRequest) (*pb.ClassifyConnectionsResponse, error) {
	// prepare request
	if len(in.GetConnections()) == 0 {
		s.logger.Warnf("service.tlsutil.classify: [peer=%s] connections(): connections is emtpy", getPeerAddr(ctx))
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	requests := make([]*tlsutil.ConnectionData, 0, len(in.GetConnections()))
	for _, r := range in.GetConnections() {
		cdata := encoding.ConnectionData(r)
		requests = append(requests, cdata)
	}
	// do request
	responses, err := s.classifier.ClassifyConnections(ctx, requests)
	if err != nil {
		s.logger.Warnf("service.tlsutil.classify: [peer=%s] connections(#%v): %v",
			getPeerAddr(ctx), len(requests), err)
		return nil, s.mapError(err)
	}
	// create response
	retResponses := make([]*pb.ClassifyConnectionsResponse_Response, 0, len(responses))
	for _, r := range responses {
		resp := &pb.ClassifyConnectionsResponse_Response{}
		retResponses = append(retResponses, resp)
		if r.Err != nil {
			resp.Err = r.Err.Error()
			continue
		}
		resp.Results = make([]*pb.ClassifyConnectionsResponse_Response_Result, 0, len(r.Results))
		for _, result := range r.Results {
			resp.Results = append(resp.Results, &pb.ClassifyConnectionsResponse_Response_Result{
				Label: result.Label,
				Prob:  result.Prob},
			)
		}
	}
	return &pb.ClassifyConnectionsResponse{Responses: retResponses}, nil
}

//mapping errors
func (s *Service) mapError(err error) error {
	switch err {
	case tlsutil.ErrCanceledRequest:
		return status.Error(codes.Canceled, err.Error())
	case tlsutil.ErrBadRequest:
		return status.Error(codes.InvalidArgument, err.Error())
	case tlsutil.ErrNotSupported:
		return status.Error(codes.Unimplemented, err.Error())
	case tlsutil.ErrUnavailable:
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Internal, tlsutil.ErrInternal.Error())
	}
}

func getPeerAddr(ctx context.Context) (paddr string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		paddr = p.Addr.String()
	}
	return
}
