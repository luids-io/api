// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package archive

import (
	"context"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/api/tlsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	logger   yalogi.Logger
	archiver tlsutil.Archiver
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
func NewService(a tlsutil.Archiver, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{archiver: a, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterArchiveServer(server, service)
}

// SaveConnection implements interface
func (s *Service) SaveConnection(ctx context.Context, req *pb.SaveConnectionRequest) (*pb.SaveConnectionResponse, error) {
	paddr := getPeerAddr(ctx)
	//parse request
	data, err := connectionFromRequest(req)
	if err != nil {
		s.logger.Warnf("invalid request from '%s': %v", paddr, err)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	//do request
	newid, err := s.archiver.SaveConnection(ctx, data)
	if err != nil {
		s.logger.Warnf("saving connection from '%s': %v", paddr, err)
		return nil, s.mapError(err)
	}
	//return response
	return &pb.SaveConnectionResponse{Id: newid}, nil
}

// SaveCertificate implements interface
func (s *Service) SaveCertificate(ctx context.Context, req *pb.SaveCertificateRequest) (*pb.SaveCertificateResponse, error) {
	paddr := getPeerAddr(ctx)
	//parse request
	data, err := certificateFromRequest(req)
	if err != nil {
		s.logger.Warnf("invalid request from '%s': %v", paddr, err)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	//do request
	newid, err := s.archiver.SaveCertificate(ctx, data)
	if err != nil {
		s.logger.Warnf("saving connection from '%s': %v", paddr, err)
		return nil, s.mapError(err)
	}
	//return response
	return &pb.SaveCertificateResponse{Id: newid}, nil
}

// StreamRecords implements interface
func (s *Service) StreamRecords(stream pb.Archive_StreamRecordsServer) error {
	paddr := getPeerAddr(stream.Context())
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		record := recordFromRequest(request)
		err = s.archiver.StoreRecord(record)
		if err != nil {
			s.logger.Warnf("saving stream record from '%s': %v", paddr, err)
			return s.mapError(err)
		}
	}
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
