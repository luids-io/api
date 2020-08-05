// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package notary

import (
	"context"
	"crypto/x509"
	"errors"
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
	// prepare request
	ip, port, sni, profile, err := parseGetServerChainRequest(in)
	if err != nil {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] getserverchain(%s,%v,%s,%s): %v",
			getPeerAddr(ctx), in.GetIp(), in.GetPort(), in.GetSni(), in.GetProfile(), err)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	// do request
	chain, err := s.notary.GetServerChain(ctx, ip, port, sni, profile)
	if err != nil {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] getserverchain(%v,%v,%s,%s): %v",
			getPeerAddr(ctx), ip, port, sni, profile, err)
		return nil, s.mapError(err)
	}
	// response
	return &pb.GetServerChainResponse{Chain: chain}, nil
}

// SetServerChain implements Service interface
func (s *Service) SetServerChain(ctx context.Context, in *pb.SetServerChainRequest) (*pb.SetServerChainResponse, error) {
	// prepare request
	ip, port, sni, profile, chain, err := parseSetServerChainRequest(in)
	if err != nil {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] setserverchain(%s,%v,%s,%s,%s): %v",
			getPeerAddr(ctx), in.GetIp(), in.GetPort(), in.GetSni(), in.GetProfile(), in.GetChain(), err)
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	// do request
	err = s.notary.SetServerChain(ctx, ip, port, sni, profile, chain)
	if err != nil {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] setserverchain(%v,%v,%s,%s,%s): %v",
			getPeerAddr(ctx), ip, port, sni, profile, chain, err)
		return nil, s.mapError(err)
	}
	// response
	return &pb.SetServerChainResponse{}, nil
}

// VerifyChain implements Service interface
func (s *Service) VerifyChain(ctx context.Context, in *pb.VerifyChainRequest) (*pb.VerifyChainResponse, error) {
	// prepare request
	chain := in.GetChain()
	if chain == "" {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] verifychain(%s,%s,%v): chain is empty",
			getPeerAddr(ctx), chain, in.GetDnsname(), in.GetForce())
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	dnsname := in.GetDnsname()
	force := in.GetForce()
	// do request
	response, err := s.notary.VerifyChain(ctx, chain, dnsname, force)
	if err != nil {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] verifychain(%s,%s,%v): %v",
			getPeerAddr(ctx), chain, in.GetDnsname(), in.GetForce(), err)
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
	// prepare request
	rawcerts := in.GetCerts()
	if len(rawcerts) == 0 {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] uploadcerts(): certs is empty", getPeerAddr(ctx))
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	certs := make([]*x509.Certificate, 0, len(rawcerts))
	for _, rawcert := range rawcerts {
		cert, err := x509.ParseCertificate(rawcert)
		if err != nil {
			s.logger.Warnf("service.tlsutil.notary: [peer=%s] uploadcerts(): %v", getPeerAddr(ctx), err)
			return nil, s.mapError(tlsutil.ErrBadRequest)
		}
		certs = append(certs, cert)
	}
	// do request
	chain, err := s.notary.UploadCerts(ctx, certs)
	if err != nil {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] uploadcerts(): %v", getPeerAddr(ctx), err)
		return nil, s.mapError(err)
	}
	// response
	return &pb.UploadCertsResponse{Chain: chain}, nil
}

// DownloadCerts implements Service interface
func (s *Service) DownloadCerts(ctx context.Context, in *pb.DownloadCertsRequest) (*pb.DownloadCertsResponse, error) {
	// prepare request
	chain := in.GetChain()
	if chain == "" {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] downloadcerts(): chain is empty", getPeerAddr(ctx))
		return nil, s.mapError(tlsutil.ErrBadRequest)
	}
	// do request
	certs, err := s.notary.DownloadCerts(ctx, chain)
	if err != nil {
		s.logger.Warnf("service.tlsutil.notary: [peer=%s] downloadcerts(%s): %v", getPeerAddr(ctx), chain, err)
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

func parseGetServerChainRequest(in *pb.GetServerChainRequest) (ip net.IP, port int, sni string, profile string, err error) {
	ip = net.ParseIP(in.GetIp())
	if ip == nil {
		err = errors.New("invalid ip")
		return
	}
	port = int(in.GetPort())
	if port <= 0 {
		err = errors.New("invalid port")
		return
	}
	sni = in.GetSni()
	profile = in.GetProfile()
	return
}

func parseSetServerChainRequest(in *pb.SetServerChainRequest) (ip net.IP, port int, sni string, profile string, chain string, err error) {
	ip = net.ParseIP(in.GetIp())
	if ip == nil {
		err = errors.New("invalid ip")
		return
	}
	port = int(in.GetPort())
	if port <= 0 {
		err = errors.New("invalid port")
		return
	}
	sni = in.GetSni()
	profile = in.GetProfile()
	chain = in.GetChain()
	if chain == "" {
		err = errors.New("chain is empty")
	}
	return
}
