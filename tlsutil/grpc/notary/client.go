// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package notary

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/api/tlsutil/grpc/pb"
	"github.com/luids-io/core/yalogi"
)

// Client provides a grpc client.
type Client struct {
	opts   clientOpts
	logger yalogi.Logger
	//grpc connection
	conn   *grpc.ClientConn
	client pb.NotaryClient
	//control
	closed bool
	//caches
	scache *serverChainCache
	ucache *uploadCache
	vcache *verifyCache
	dcache *downloadCache
}

// ClientOption encapsules options for client.
type ClientOption func(*clientOpts)

type clientOpts struct {
	logger    yalogi.Logger
	closeConn bool
	//cache opts
	useCache         bool
	ttl, negativettl int
	cacheCleanup     time.Duration
}

var defaultClientOpts = clientOpts{
	logger:       yalogi.LogNull,
	closeConn:    true,
	cacheCleanup: defaultCacheCleanups,
}

// CloseConnection option closes grpc connection on close.
func CloseConnection(b bool) ClientOption {
	return func(o *clientOpts) {
		o.closeConn = b
	}
}

// SetLogger option allows set a custom logger.
func SetLogger(l yalogi.Logger) ClientOption {
	return func(o *clientOpts) {
		if l != nil {
			o.logger = l
		}
	}
}

// SetCache sets cache ttl and negative ttl (for validations).
func SetCache(ttl, negativettl int) ClientOption {
	return func(o *clientOpts) {
		if ttl > 0 {
			o.ttl = ttl
			o.negativettl = negativettl
			o.useCache = true
		}
	}
}

// SetCacheCleanUps sets interval between cache cleanups.
func SetCacheCleanUps(d time.Duration) ClientOption {
	return func(o *clientOpts) {
		if d > 0 {
			o.cacheCleanup = d
		}
	}
}

// NewClient returns a new client.
func NewClient(conn *grpc.ClientConn, opt ...ClientOption) *Client {
	opts := defaultClientOpts
	for _, o := range opt {
		o(&opts)
	}
	c := &Client{
		opts:   opts,
		logger: opts.logger,
		conn:   conn,
		client: pb.NewNotaryClient(conn),
	}
	if c.opts.useCache {
		c.scache = newServerChainCache(c.opts.ttl, c.opts.cacheCleanup)
		c.ucache = newUploadCache(c.opts.ttl, c.opts.cacheCleanup)
		c.vcache = newVerifyCache(c.opts.ttl, c.opts.negativettl, c.opts.cacheCleanup)
		c.dcache = newDownloadCache(c.opts.ttl, c.opts.cacheCleanup)
	}
	return c
}

// GetServerChain implements tlsutil.Notary interface.
func (c *Client) GetServerChain(ctx context.Context, ip net.IP, port int, sni, profile string) (string, error) {
	if c.closed {
		c.logger.Warnf("client.tlsutil.notary: getserverchain(%v,%v,%s,%s): client is closed",
			ip, port, sni, profile)
		return "", tlsutil.ErrUnavailable
	}
	if c.opts.useCache {
		chain, ok := c.scache.get(ip, port, sni, profile)
		if ok {
			return chain, nil
		}
	}
	//prepare request
	req := &pb.GetServerChainRequest{
		Ip:      ip.String(),
		Port:    int32(port),
		Sni:     sni,
		Profile: profile,
	}
	//do request
	resp, err := c.client.GetServerChain(ctx, req)
	if err != nil {
		c.logger.Warnf("client.tlsutil.notary: getserverchain(%v,%v,%s,%s): %v",
			ip, port, sni, profile, err)
		return "", c.mapError(err)
	}
	//get response
	chain := resp.GetChain()
	if c.opts.useCache {
		c.scache.set(ip, port, sni, profile, chain)
	}
	return chain, nil
}

// SetServerChain implements tlsutil.Notary interface.
func (c *Client) SetServerChain(ctx context.Context, ip net.IP, port int, sni, profile string, chain string) error {
	if c.closed {
		c.logger.Warnf("client.tlsutil.notary: setserverchain(%v,%v,%s,%s,%s): client is closed",
			ip, port, sni, profile, chain)
		return tlsutil.ErrUnavailable
	}
	if c.opts.useCache {
		//check if it's the same cached chain
		cached, ok := c.scache.get(ip, port, sni, profile)
		if ok && cached == chain {
			//only updates cache
			c.scache.set(ip, port, sni, profile, chain)
			return nil
		}
	}
	//prepare request
	req := &pb.SetServerChainRequest{
		Ip:      ip.String(),
		Port:    int32(port),
		Sni:     sni,
		Profile: profile,
		Chain:   chain,
	}
	//do request
	_, err := c.client.SetServerChain(ctx, req)
	if err != nil {
		c.logger.Warnf("client.tlsutil.notary: setserverchain(%v,%v,%s,%s): %v",
			ip, port, sni, profile, chain, err)
		return c.mapError(err)
	}
	if c.opts.useCache {
		c.scache.set(ip, port, sni, profile, chain)
	}
	return nil
}

// VerifyChain implements tlsutil.Notary interface.
func (c *Client) VerifyChain(ctx context.Context, chain string, dnsname string, force bool) (tlsutil.VerifyResponse, error) {
	if c.closed {
		c.logger.Warnf("client.tlsutil.notary: verifychain(%s,%s,%v): client is closed", chain, dnsname, force)
		return tlsutil.VerifyResponse{}, tlsutil.ErrUnavailable
	}
	if c.opts.useCache && !force {
		resp, ok := c.vcache.get(chain, dnsname)
		if ok {
			return resp, nil
		}
	}
	//prepare request
	req := &pb.VerifyChainRequest{
		Chain:   chain,
		Dnsname: dnsname,
		Force:   force,
	}
	//do verify
	res, err := c.client.VerifyChain(ctx, req)
	if err != nil {
		c.logger.Warnf("client.tlsutil.notary: verifychain(%s,%s,%v): %v",
			chain, dnsname, force, err)
		return tlsutil.VerifyResponse{}, c.mapError(err)
	}
	//get response
	vr := tlsutil.VerifyResponse{
		Invalid: res.GetInvalid(),
		Reason:  res.GetReason(),
		TTL:     int(res.GetTTL()),
	}
	if c.opts.useCache {
		vr = c.vcache.set(chain, dnsname, vr)
	}
	return vr, nil
}

// UploadCerts implements tlsutil.Notary interface.
func (c *Client) UploadCerts(ctx context.Context, certs []*x509.Certificate) (string, error) {
	if c.closed {
		c.logger.Warnf("client.tlsutil.notary: uploadcerts(): client is closed")
		return "", tlsutil.ErrUnavailable
	}
	var chain, cachekey string
	if c.opts.useCache {
		var ok bool
		chain, cachekey, ok = c.ucache.get(certs)
		if ok {
			return chain, nil
		}
	}
	//prepare request
	rawcerts := make([][]byte, 0, len(certs))
	for _, cert := range certs {
		rawcerts = append(rawcerts, cert.Raw)
	}
	req := &pb.UploadCertsRequest{Certs: rawcerts}
	//do upload
	res, err := c.client.UploadCerts(ctx, req)
	if err != nil {
		c.logger.Warnf("client.tlsutil.notary: uploadcerts(): %v", err)
		return "", c.mapError(err)
	}
	//return response
	chain = res.GetChain()
	if c.opts.useCache {
		c.ucache.set(cachekey, chain)
	}
	return chain, nil
}

// DownloadCerts implements tlsutil.Notary interface.
func (c *Client) DownloadCerts(ctx context.Context, chain string) ([]*x509.Certificate, error) {
	if c.closed {
		c.logger.Warnf("client.tlsutil.notary: downloadcerts(%s): client is closed", chain)
		return nil, tlsutil.ErrUnavailable
	}
	if c.opts.useCache {
		certs, ok := c.dcache.get(chain)
		if ok {
			return certs, nil
		}
	}
	//prepare request
	req := &pb.DownloadCertsRequest{Chain: chain}
	//do download
	res, err := c.client.DownloadCerts(ctx, req)
	if err != nil {
		c.logger.Warnf("client.tlsutil.notary: downloadcerts(%s): %v", chain, err)
		return nil, c.mapError(err)
	}
	//process response
	var certs []*x509.Certificate
	rawcerts := res.GetCerts()
	if len(rawcerts) > 0 {
		certs = make([]*x509.Certificate, 0, len(rawcerts))
		for _, rawcert := range rawcerts {
			cert, err := x509.ParseCertificate(rawcert)
			if err != nil {
				c.logger.Warnf("client.tlsutil.notary: downloadcerts(%s): parsing cert: %v", chain, err)
				return nil, tlsutil.ErrInternal
			}
			certs = append(certs, cert)
		}
	}
	if c.opts.useCache {
		c.dcache.set(chain, certs)
	}
	return certs, nil
}

// Ping checks connectivity with the api.
func (c *Client) Ping() error {
	if c.closed {
		return errors.New("client closed")
	}
	st := c.conn.GetState()
	switch st {
	case connectivity.TransientFailure:
		return fmt.Errorf("connection state: %v", st)
	case connectivity.Shutdown:
		return fmt.Errorf("connection state: %v", st)
	}
	return nil
}

//Close the client.
func (c *Client) Close() error {
	if c.closed {
		return errors.New("client closed")
	}
	c.closed = true
	if c.opts.closeConn {
		return c.conn.Close()
	}
	return nil
}

//mapping errors
func (c *Client) mapError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.InvalidArgument:
		return tlsutil.ErrBadRequest
	case codes.Unimplemented:
		return tlsutil.ErrNotSupported
	case codes.Internal:
		return tlsutil.ErrInternal
	case codes.Unavailable:
		if st.Message() == tlsutil.ErrDialingWithServer.Error() {
			return tlsutil.ErrDialingWithServer
		}
		return tlsutil.ErrUnavailable
	case codes.NotFound:
		return tlsutil.ErrChainNotFound
	default:
		return tlsutil.ErrUnavailable
	}
}

//API returns API service name implemented.
func (c *Client) API() string {
	return ServiceName()
}
