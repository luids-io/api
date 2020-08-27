// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/netutil"
	"github.com/luids-io/api/netutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a grpc service wrapper.
type Service struct {
	logger     yalogi.Logger
	expiration time.Duration
	factory    netutil.AnalyzerFactory
}

// ServiceOption is used for service configuration
type ServiceOption func(*serviceOpts)

type serviceOpts struct {
	logger     yalogi.Logger
	expiration time.Duration
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

// SetPacketExpiration option allows set packet expire time.
func SetPacketExpiration(d time.Duration) ServiceOption {
	return func(o *serviceOpts) {
		if d > 0 {
			o.expiration = d
		}
	}
}

// NewService returns a new Service.
func NewService(f netutil.AnalyzerFactory, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{factory: f, logger: opts.logger, expiration: opts.expiration}
}

// RegisterServer registers a service in the grpc server.
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterAnalyzeServer(server, service)
}

// SendPackets implements grpc api.
func (s *Service) SendPackets(stream pb.Analyze_SendPacketsServer) error {
	paddr := getPeerAddr(stream.Context())
	if paddr == "" {
		s.logger.Errorf("service.netutil.analyze: sendpackets(): can't get peer address")
		return status.Errorf(codes.Internal, netutil.ErrInternal.Error())
	}
	analyzer, err := s.factory.NewAnalyzer(paddr)
	if err != nil {
		s.logger.Warnf("service.netutil.analyze: [peer=%s] sendpackets(): %v", paddr, err)
		return s.mapError(err)
	}
	defer analyzer.Close()

	var lastTs time.Time
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			s.logger.Warnf("service.netutil.analyze: [peer=%s] sendpackets(): error receiving: %v", paddr, err)
			return err
		}
		// parse request
		layer, md, pdata, err := parseSendPacketRequest(req)
		if err != nil {
			s.logger.Warnf("service.netutil.analyze: [peer=%s] sendpackets(): %v", paddr, err)
			return s.mapError(netutil.ErrBadRequest)
		}
		// check timestamp...
		now := time.Now()
		ts := md.Timestamp
		if ts.IsZero() || err != nil {
			s.logger.Warnf("service.netutil.analyze: [peer=%s] sendpackets(): invalid timestamp", paddr)
			return s.mapError(netutil.ErrBadRequest)
		}
		//check if out of order
		if !lastTs.IsZero() && lastTs.After(ts) {
			err := netutil.ErrPacketOutOfOrder
			s.logger.Warnf("service.netutil.analyze: [peer=%s] sendpackets(): %v: %v>%v", paddr, err, ts, lastTs)
			return s.mapError(err)
		}
		lastTs = ts
		//check if out of sync
		if s.expiration > 0 {
			if ts.After(now) {
				if ts.Sub(now) > s.expiration {
					err := netutil.ErrTimeOutOfSync
					s.logger.Warnf("service.netutil.analyze: [peer=%s] sendpackets(): %v: %v>%v ", paddr, err, now, ts)
					return s.mapError(err)
				}
			} else {
				if now.Sub(ts) > s.expiration {
					err := netutil.ErrTimeOutOfSync
					s.logger.Warnf("service.netutil.analyze: [peer=%s] sendpackets(): %v: %v<%v ", paddr, err, now, ts)
					return s.mapError(err)
				}
			}
		}
		// send message
		err = analyzer.SendPacket(layer, md, pdata)
		if err != nil {
			s.logger.Warnf("service.netutil.analyze: [peer=%s] sendpackets(layer=[%v],md=[%v]): %v", paddr, layer, md, err)
			return s.mapError(err)
		}
	}
}

//mapping errors
func (s *Service) mapError(err error) error {
	switch err {
	case netutil.ErrBadRequest:
		return status.Error(codes.InvalidArgument, err.Error())
	case netutil.ErrPacketOutOfOrder:
		return status.Error(codes.OutOfRange, err.Error())
	case netutil.ErrTimeOutOfSync:
		return status.Error(codes.OutOfRange, err.Error())
	case netutil.ErrNotSupported:
		return status.Error(codes.Unimplemented, err.Error())
	case netutil.ErrAnalyzerExists:
		return status.Error(codes.ResourceExhausted, err.Error())
	case netutil.ErrUnavailable:
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Internal, netutil.ErrInternal.Error())
	}
}

func parseSendPacketRequest(req *pb.SendPacketRequest) (layer netutil.Layer, md netutil.PacketMetadata, pdata []byte, err error) {
	// get layer
	layer = netutil.Layer(req.GetLayer())
	if layer < netutil.Ethernet || layer > netutil.IPv6 {
		err = errors.New("invalid layer")
		return
	}
	// get metadata
	reqMD := req.GetMetadata()
	if reqMD == nil {
		err = errors.New("no metadata")
		return
	}
	md.Timestamp, err = ptypes.Timestamp(reqMD.Timestamp)
	md.CaptureLength = int(reqMD.GetCaptureLength())
	md.Length = int(reqMD.GetLength())
	md.InterfaceIndex = int(reqMD.GetInterfaceIndex())
	// get data
	pdata = req.GetData()
	if pdata == nil {
		err = errors.New("no packet data")
	}
	return
}

func getPeerAddr(ctx context.Context) (paddr string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		paddr = p.Addr.String()
	}
	return
}
