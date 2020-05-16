// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"context"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/api/tlsutil/grpc/encoding"
	"github.com/luids-io/api/tlsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	logger  yalogi.Logger
	factory tlsutil.AnalyzerFactory
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
func NewService(f tlsutil.AnalyzerFactory, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{factory: f, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterAnalyzeServer(server, service)
}

// SendMessages manage messages
func (s *Service) SendMessages(stream pb.Analyze_SendMessagesServer) error {
	paddr := getPeerAddr(stream.Context())
	if paddr == "" {
		s.logger.Errorf("can't get peer address")
		return status.Errorf(codes.Internal, tlsutil.ErrInternal.Error())
	}
	analyzer, err := s.factory.NewAnalyzer(paddr)
	if err != nil {
		return s.mapError(err)
	}
	defer analyzer.Close()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			s.logger.Warnf("receiving from '%s': %v", paddr, err)
			return err
		}
		msg := encoding.MessageRequest(req)
		err = analyzer.SendMessage(msg)
		if err != nil {
			s.logger.Warnf("analyzing from '%s' msg=[%v]: %v", paddr, msg, err)
			if err != tlsutil.ErrStreamNotFound {
				return s.mapError(err)
			}
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
	case tlsutil.ErrDuplicatedStream:
		return status.Error(codes.AlreadyExists, tlsutil.ErrDuplicatedStream.Error())
	case tlsutil.ErrStreamNotFound:
		return status.Error(codes.FailedPrecondition, tlsutil.ErrStreamNotFound.Error())
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
