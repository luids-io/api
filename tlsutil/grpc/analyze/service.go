// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"context"
	"io"
	"time"

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
	logger     yalogi.Logger
	expiration time.Duration
	factory    tlsutil.AnalyzerFactory
}

type serviceOpts struct {
	logger     yalogi.Logger
	expiration time.Duration
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

// SetPacketExpiration option allows set a custom logger
func SetPacketExpiration(d time.Duration) ServiceOption {
	return func(o *serviceOpts) {
		if d > 0 {
			o.expiration = d
		}
	}
}

// NewService returns a new Service for the grpc api
func NewService(f tlsutil.AnalyzerFactory, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{factory: f, logger: opts.logger, expiration: opts.expiration}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterAnalyzeServer(server, service)
}

// SendMessages manage messages
func (s *Service) SendMessages(stream pb.Analyze_SendMessagesServer) error {
	paddr := getPeerAddr(stream.Context())
	if paddr == "" {
		s.logger.Errorf("service.tlsutil.analyze: sendmessages(): can't get peer address")
		return status.Errorf(codes.Internal, tlsutil.ErrInternal.Error())
	}
	analyzer, err := s.factory.NewAnalyzer(paddr)
	if err != nil {
		s.logger.Warnf("service.tlsutil.analyze: [peer=%s] sendmessages(): %v", paddr, err)
		return s.mapError(err)
	}
	defer analyzer.Close()

	var lastTs time.Time
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			s.logger.Warnf("service.tlsutil.analyze: [peer=%s] sendmessages(): error receiving: %v", paddr, err)
			return err
		}
		msg := encoding.MessageRequest(req)
		// check timestamp
		if msg.Data != nil {
			now := time.Now()
			ts := msg.Data.Timestamp
			//check if out of order
			if !lastTs.IsZero() && lastTs.After(ts) {
				err := tlsutil.ErrMsgOutOfOrder
				s.logger.Warnf("service.tlsutil.analyze: [peer=%s] sendmessages(): %v: %v>%v", paddr, err, ts, lastTs)
				return s.mapError(err)
			}
			lastTs = ts
			//check if expires
			if s.expiration > 0 {
				if ts.After(now) {
					if ts.Sub(now) > s.expiration {
						err := tlsutil.ErrTimeOutOfSync
						s.logger.Warnf("service.tlsutil.analyze: [peer=%s] sendmessages(): %v: %v>%v ", paddr, err, now, ts)
						return s.mapError(err)
					}
				} else {
					if now.Sub(ts) > s.expiration {
						err := tlsutil.ErrTimeOutOfSync
						s.logger.Warnf("service.tlsutil.analyze: [peer=%s] sendmessages(): %v: %v<%v ", paddr, err, now, ts)
						return s.mapError(err)
					}
				}
			}
		}
		// send message
		err = analyzer.SendMessage(msg)
		if err != nil {
			s.logger.Warnf("service.tlsutil.analyze: [peer=%s] sendmessages([msg=%v]): %v", paddr, msg, err)
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
	case tlsutil.ErrMsgOutOfOrder:
		return status.Error(codes.OutOfRange, err.Error())
	case tlsutil.ErrTimeOutOfSync:
		return status.Error(codes.OutOfRange, err.Error())
	case tlsutil.ErrNotSupported:
		return status.Error(codes.Unimplemented, err.Error())
	case tlsutil.ErrUnavailable:
		return status.Error(codes.Unavailable, err.Error())
	case tlsutil.ErrAnalyzerExists:
		return status.Error(codes.ResourceExhausted, err.Error())
	case tlsutil.ErrDuplicatedStream:
		return status.Error(codes.AlreadyExists, err.Error())
	case tlsutil.ErrStreamNotFound:
		return status.Error(codes.FailedPrecondition, err.Error())
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
