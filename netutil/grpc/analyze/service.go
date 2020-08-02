// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"context"
	"io"
	"time"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/netutil"
	"github.com/luids-io/api/netutil/grpc/pb"
	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/core/yalogi"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	logger     yalogi.Logger
	expiration time.Duration
	factory    netutil.AnalyzerFactory
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
func NewService(f netutil.AnalyzerFactory, opt ...ServiceOption) *Service {
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

// SendPackets manage messages
func (s *Service) SendPackets(stream pb.Analyze_SendPacketsServer) error {
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
		// get layer
		reqLayer := req.GetLayer()
		if reqLayer < 0 || reqLayer > pb.Layer_IPV6 {
			s.logger.Warnf("receiving from '%s': invalid layer", paddr)
			return s.mapError(tlsutil.ErrBadRequest)
		}
		layer := netutil.Layer(reqLayer)

		// get metadata
		reqMD := req.GetMetadata()
		if reqMD == nil {
			s.logger.Warnf("receiving from '%s': no metadata", paddr)
			return s.mapError(tlsutil.ErrBadRequest)
		}
		// check timestamp
		now := time.Now()
		ts, _ := ptypes.Timestamp(reqMD.Timestamp)
		if ts.IsZero() {
			s.logger.Warnf("receiving from '%s': invalid timestamp", paddr)
			return s.mapError(tlsutil.ErrBadRequest)
		}
		if s.expiration > 0 {
			if ts.After(now) {
				if ts.Sub(now) > s.expiration {
					s.logger.Warnf("receiving from '%s': out of sync", paddr)
					return s.mapError(tlsutil.ErrTimeOutOfSync)
				}
			} else {
				if now.Sub(ts) > s.expiration {
					s.logger.Warnf("receiving from '%s': out of sync", paddr)
					return s.mapError(tlsutil.ErrTimeOutOfSync)
				}
			}
		}
		md := netutil.PacketMetadata{}
		md.Timestamp = ts
		md.CaptureLength = int(reqMD.GetCaptureLength())
		md.Length = int(reqMD.GetLength())
		md.InterfaceIndex = int(reqMD.GetInterfaceIndex())

		// send message
		err = analyzer.SendPacket(layer, req.GetData(), md)
		if err != nil {
			s.logger.Warnf("analyzing from '%s' layer=[%v],ts=[%v],len(data)=[%v]: %v", paddr, layer, ts, len(req.GetData()), err)
			if err != tlsutil.ErrStreamNotFound {
				return s.mapError(err)
			}
		}
	}
}

//mapping errors
func (s *Service) mapError(err error) error {
	switch err {
	case netutil.ErrBadRequest:
		return status.Error(codes.InvalidArgument, err.Error())
	case netutil.ErrTimeOutOfSync:
		return status.Error(codes.OutOfRange, err.Error())
	case netutil.ErrNotSupported:
		return status.Error(codes.Unimplemented, err.Error())
	case netutil.ErrUnavailable:
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Internal, netutil.ErrInternal.Error())
	}
}

func getPeerAddr(ctx context.Context) (paddr string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		paddr = p.Addr.String()

	}
	return
}
