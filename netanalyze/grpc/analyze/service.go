// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package analyze

import (
	"fmt"
	"io"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/netanalyze"
	"github.com/luids-io/api/netanalyze/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a service wrapper for the grpc api
type Service struct {
	logger     yalogi.Logger
	expiration time.Duration
	pcktsvc    netanalyze.Service
	ethproc    netanalyze.Processor
	ip4proc    netanalyze.Processor
	ip6proc    netanalyze.Processor
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
func NewService(p netanalyze.Service, ethproc, ip4proc, ip6proc netanalyze.Processor, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{
		logger:     opts.logger,
		expiration: opts.expiration,
		pcktsvc:    p,
		ethproc:    ethproc,
		ip4proc:    ip4proc,
		ip6proc:    ip6proc,
	}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterAnalyzeServer(server, service)
}

// SendEtherPackets manage requests
func (s *Service) SendEtherPackets(stream pb.Analyze_SendEtherPacketsServer) error {
	if s.ethproc == nil {
		return status.Errorf(codes.Unavailable, netanalyze.ErrUnavailable.Error())
	}
	return s.sendPackets(stream, layers.LinkTypeEthernet)
}

// SendIPv4Packets manage requests
func (s *Service) SendIPv4Packets(stream pb.Analyze_SendIPv4PacketsServer) error {
	if s.ip4proc == nil {
		return status.Errorf(codes.Unavailable, netanalyze.ErrUnavailable.Error())
	}
	return s.sendPackets(stream, layers.LinkTypeIPv4)
}

// SendIPv6Packets manage requests
func (s *Service) SendIPv6Packets(stream pb.Analyze_SendIPv6PacketsServer) error {
	if s.ip6proc == nil {
		return status.Errorf(codes.Unavailable, netanalyze.ErrUnavailable.Error())
	}
	return s.sendPackets(stream, layers.LinkTypeIPv6)
}

// sendPackets manage requests
func (s *Service) sendPackets(stream pcktServerStream, linkType layers.LinkType) error {
	ctx := stream.Context()
	p, ok := peer.FromContext(ctx)
	if !ok {
		s.logger.Errorf("can't get peer address")
		return status.Errorf(codes.Internal, netanalyze.ErrInternal.Error())
	}
	// creates packet source
	name := p.Addr.String()
	var proc netanalyze.Processor
	switch linkType {
	case layers.LinkTypeEthernet:
		name = fmt.Sprintf("%s-eth", name)
		proc = s.ethproc
	case layers.LinkTypeIPv4:
		name = fmt.Sprintf("%s-ip4", name)
		proc = s.ip4proc
	case layers.LinkTypeIPv6:
		name = fmt.Sprintf("%s-ip6", name)
		proc = s.ip6proc
	default:
		s.logger.Errorf("invalid send layer from '%s'", name)
		return status.Errorf(codes.Internal, netanalyze.ErrInternal.Error())
	}
	psource := &pcktSource{
		expiration: s.expiration,
		err:        make(chan error),
		stream:     stream,
	}
	err := s.pcktsvc.Register(name, psource, proc)
	if err != nil {
		s.logger.Warnf("registering '%s': %v", name, err)
		return status.Errorf(codes.Internal, netanalyze.ErrInternal.Error())
	}
	//waits for close or error
	err = <-psource.err
	//clean
	s.pcktsvc.Unregister(name)
	close(psource.err)
	if err == io.EOF {
		return nil
	}
	s.logger.Warnf("processing '%s': %v", name, err)
	if err == netanalyze.ErrTimeOutOfSync {
		return status.Errorf(codes.OutOfRange, netanalyze.ErrTimeOutOfSync.Error())
	}
	return status.Errorf(codes.Internal, netanalyze.ErrInternal.Error())
}

type pcktSource struct {
	expiration time.Duration
	stream     pcktServerStream
	err        chan error
	closed     bool
}

type pcktServerStream interface {
	grpc.ServerStream
	Recv() (*pb.SendPacketRequest, error)
}

func (p *pcktSource) ReadPacketData() (data []byte, ci gopacket.CaptureInfo, err error) {
	if p.closed {
		err = io.EOF
		return
	}
	var req *pb.SendPacketRequest
	req, err = p.stream.Recv()
	if err != nil {
		p.closed = true
		p.err <- err
		return
	}
	meta := req.GetMetadata()
	data = req.GetData()
	// check timestamp
	now := time.Now()
	ts, _ := ptypes.Timestamp(meta.GetTimestamp())
	if p.expiration > 0 {
		if ts.After(now) {
			if ts.Sub(now) > p.expiration {
				p.closed = true
				p.err <- netanalyze.ErrTimeOutOfSync
				return
			}
		} else {
			if now.Sub(ts) > p.expiration {
				p.closed = true
				p.err <- netanalyze.ErrTimeOutOfSync
				return
			}
		}
	}
	ci.Timestamp = ts
	ci.InterfaceIndex = int(meta.GetInterfaceIndex())
	ci.CaptureLength = len(data)
	ci.Length = len(data)
	return
}
