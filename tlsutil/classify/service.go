// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package classify

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/luids-io/api/protogen/tlsutilpb"
	"github.com/luids-io/api/tlsutil/encoding"
	"github.com/luids-io/core/tlsutil"
)

// Service provides a grpc wrapper
type Service struct {
	opts       serviceOpts
	classifier tlsutil.Classifier
}

type serviceOpts struct {
	disclosureErr bool
	dataBuff      int
}

var defaultServiceOpts = serviceOpts{dataBuff: 100}

// ServiceOption is used for service configuration
type ServiceOption func(*serviceOpts)

// DisclosureErrors returns errors without replacing by a generic message
func DisclosureErrors(b bool) ServiceOption {
	return func(o *serviceOpts) {
		o.disclosureErr = b
	}
}

// SetDataBuff option allows change channel buffer data
func SetDataBuff(i int) ServiceOption {
	return func(o *serviceOpts) {
		o.dataBuff = i
	}
}

// NewService returns a new Service for the cheker
func NewService(c tlsutil.Classifier, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{classifier: c, opts: opts}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterClassifyServer(server, service)
}

// Connections implements grpc interface
func (s *Service) Connections(ctx context.Context, in *pb.ClassifyConnectionsRequest) (*pb.ClassifyConnectionsResponse, error) {
	// prepare request
	if len(in.GetConnections()) == 0 {
		rpcerr := status.Error(codes.InvalidArgument, "connections is empty")
		return nil, rpcerr
	}
	requests := make([]*tlsutil.ConnectionData, 0, len(in.GetConnections()))
	for _, r := range in.GetConnections() {
		cdata := encoding.ConnectionData(r)
		requests = append(requests, cdata)
	}
	// do request
	responses, err := s.classifier.ClassifyConnections(ctx, requests)
	if err != nil {
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
	//TODO
	rpcerr := status.Error(codes.Unavailable, "service not available")
	if s.opts.disclosureErr {
		rpcerr = status.Error(codes.Unavailable, err.Error())
	}
	return rpcerr
}
