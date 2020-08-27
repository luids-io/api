// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package archive

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/dnsutil"
	"github.com/luids-io/api/dnsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service implements a grpc service wrapper.
type Service struct {
	logger   yalogi.Logger
	archiver dnsutil.Archiver
}

// ServiceOption is used for service configuration.
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

// NewService returns a new Service.
func NewService(a dnsutil.Archiver, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{archiver: a, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server.
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterArchiveServer(server, service)
}

// SaveResolv implements grpc interface.
func (s *Service) SaveResolv(ctx context.Context, req *pb.SaveResolvRequest) (*pb.SaveResolvResponse, error) {
	//parse request
	data, err := parseRequest(req)
	if err != nil {
		s.logger.Warnf("service.dnsutil.archive: [peer=%s] saveresolv(%s,%v): %v", getPeerAddr(ctx), data.Name, data.Client, err)
		return nil, s.mapError(dnsutil.ErrBadRequest)
	}
	//do request
	newid, err := s.archiver.SaveResolv(ctx, data)
	if err != nil {
		s.logger.Warnf("service.dnsutil.archive: [peer=%s] saveresolv(%s,%v): %v", getPeerAddr(ctx), data.Name, data.Client, err)
		return nil, s.mapError(err)
	}
	//return response
	return &pb.SaveResolvResponse{Id: newid}, nil
}

func parseRequest(req *pb.SaveResolvRequest) (dnsutil.ResolvData, error) {
	i := dnsutil.ResolvData{}
	i.Timestamp, _ = ptypes.Timestamp(req.GetTs())
	i.Duration = time.Duration(req.GetDuration())
	i.Server = net.ParseIP(req.GetServerIp())
	if i.Server == nil {
		return i, errors.New("bad server ip")
	}
	i.Client = net.ParseIP(req.GetClientIp())
	if i.Client == nil {
		return i, errors.New("bad client ip")
	}
	i.QID = uint16(req.GetQid())
	i.Name = req.GetName()
	i.CheckingDisabled = req.GetCheckingDisabled()
	i.ReturnCode = int(req.GetReturnCode())
	i.AuthenticatedData = req.GetAuthenticatedData()
	if len(req.GetResolvedIps()) > 0 {
		i.Resolved = make([]net.IP, 0, len(req.GetResolvedIps()))
		for _, r := range req.GetResolvedIps() {
			ip := net.ParseIP(r)
			if ip == nil {
				return i, errors.New("bad resolved ip")
			}
			i.Resolved = append(i.Resolved, ip)
		}
	}
	return i, nil
}

//mapping errors
func (s *Service) mapError(err error) error {
	switch err {
	case dnsutil.ErrCanceledRequest:
		return status.Error(codes.Canceled, err.Error())
	case dnsutil.ErrBadRequest:
		return status.Error(codes.InvalidArgument, err.Error())
	case dnsutil.ErrNotSupported:
		return status.Error(codes.Unimplemented, err.Error())
	case dnsutil.ErrUnavailable:
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Internal, dnsutil.ErrInternal.Error())
	}
}

func getPeerAddr(ctx context.Context) (paddr string) {
	p, ok := peer.FromContext(ctx)
	if ok {
		paddr = p.Addr.String()
	}
	return
}
