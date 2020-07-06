// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package notary

import (
	"context"
	"crypto/x509"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/api/tlsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Service provides a wrapper for the interface tlsutil.Notary that handles
// grpc requests.
type Service struct {
	logger yalogi.Logger
	notary tlsutil.Notary
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

// NewService returns a new Service for the notary
func NewService(notary tlsutil.Notary, opt ...ServiceOption) *Service {
	opts := defaultServiceOpts
	for _, o := range opt {
		o(&opts)
	}
	return &Service{notary: notary, logger: opts.logger}
}

// RegisterServer registers a service in the grpc server
func RegisterServer(server *grpc.Server, service *Service) {
	pb.RegisterNotaryServer(server, service)
}

// GetServerChain implements Service interface
func (s *Service) GetServerChain(ctx context.Context, in *pb.GetServerChainRequest) (*pb.GetServerChainResponse, error) {
	paddr := getPeerAddr(ctx)
	// prepare request
	ip := net.ParseIP(in.GetIp())
	if ip == nil {
		s.logger.Warnf("get server chain '%s': ip invalid", paddr)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	port := int(in.GetPort())
	if port <= 0 {
		s.logger.Warnf("get server chain '%s': port invalid", paddr)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	sni := in.GetSni()
	profile := in.GetProfile()
	// do request
	chain, err := s.notary.GetServerChain(ctx, ip, port, sni, profile)
	if err != nil {
		s.logger.Warnf("get server chain '%s': %v", paddr, err)
		return nil, s.mapError(err)
	}
	// response
	return &pb.GetServerChainResponse{Chain: chain}, nil
}

// SetServerChain implements Service interface
func (s *Service) SetServerChain(ctx context.Context, in *pb.SetServerChainRequest) (*pb.SetServerChainResponse, error) {
	paddr := getPeerAddr(ctx)
	// prepare request
	ip := net.ParseIP(in.GetIp())
	if ip == nil {
		s.logger.Warnf("set server chain '%s': ip invalid", paddr)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	port := int(in.GetPort())
	if port <= 0 {
		s.logger.Warnf("set server chain '%s': port invalid", paddr)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	sni := in.GetSni()
	profile := in.GetProfile()
	chain := in.GetChain()
	if chain == "" {
		s.logger.Warnf("set server chain '%s': chain is empty", paddr)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	// do request
	err := s.notary.SetServerChain(ctx, ip, port, sni, profile, chain)
	if err != nil {
		s.logger.Warnf("set server chain '%s': %v", paddr, err)
		return nil, s.mapError(err)
	}
	// response
	return &pb.SetServerChainResponse{}, nil
}

// VerifyChain implements Service interface
func (s *Service) VerifyChain(ctx context.Context, in *pb.VerifyChainRequest) (*pb.VerifyChainResponse, error) {
	paddr := getPeerAddr(ctx)
	// prepare request
	chain := in.GetChain()
	if chain == "" {
		s.logger.Warnf("verify chain '%s': chain is empty", paddr)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	dnsname := in.GetDnsname()
	force := in.GetForce()
	// do request
	response, err := s.notary.VerifyChain(ctx, chain, dnsname, force)
	if err != nil {
		s.logger.Warnf("verify chain '%s': %v", paddr, err)
		return nil, s.mapError(err)
	}
	// response
	res := &pb.VerifyChainResponse{
		Invalid: response.Invalid,
		Reason:  response.Reason,
		TTL:     int32(response.TTL),
	}
	return res, nil
}

// UploadCerts implements Service interface
func (s *Service) UploadCerts(ctx context.Context, in *pb.UploadCertsRequest) (*pb.UploadCertsResponse, error) {
	paddr := getPeerAddr(ctx)
	// prepare request
	rawcerts := in.GetCerts()
	if len(rawcerts) == 0 {
		s.logger.Warnf("upload certs '%s': certs is empty", paddr)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	certs := make([]*x509.Certificate, 0, len(rawcerts))
	for _, rawcert := range rawcerts {
		cert, err := x509.ParseCertificate(rawcert)
		if err != nil {
			s.logger.Warnf("upload certs '%s': %v", paddr, err)
			return nil, s.mapError(tlsutil.ErrBadRequest)
		}
		certs = append(certs, cert)
	}
	// do request
	chain, err := s.notary.UploadCerts(ctx, certs)
	if err != nil {
		s.logger.Warnf("upload certs '%s': %v", paddr, err)
		return nil, s.mapError(err)
	}
	// response
	return &pb.UploadCertsResponse{Chain: chain}, nil
}

// DownloadCerts implements Service interface
func (s *Service) DownloadCerts(ctx context.Context, in *pb.DownloadCertsRequest) (*pb.DownloadCertsResponse, error) {
	paddr := getPeerAddr(ctx)
	// prepare request
	chain := in.GetChain()
	if chain == "" {
		s.logger.Warnf("download certs '%s': chain is empty", paddr)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	// do request
	certs, err := s.notary.DownloadCerts(ctx, chain)
	if err != nil {
		s.logger.Warnf("download certs '%s': %v", paddr, err)
		return nil, s.mapError(err)
	}
	// response
	var rawcerts [][]byte
	if len(certs) > 0 {
		rawcerts = make([][]byte, 0, len(certs))
		for _, cert := range certs {
			rawcerts = append(rawcerts, cert.Raw)
		}
	}
	return &pb.DownloadCertsResponse{Certs: rawcerts}, nil
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
	case tlsutil.ErrChainNotFound:
		return status.Error(codes.NotFound, err.Error())
	case tlsutil.ErrDialingWithServer:
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
