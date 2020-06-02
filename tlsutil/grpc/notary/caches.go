// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package notary

import (
	"crypto/sha1"
	"crypto/x509"
	"fmt"
	"net"
	"time"

	"github.com/luids-io/api/tlsutil"
	"github.com/luids-io/core/certverify"
	cacheimpl "github.com/patrickmn/go-cache"
)

const defaultCacheCleanups = 3 * time.Minute

type serverChainCache struct {
	cache *cacheimpl.Cache
}

func newServerChainCache(ttl, cleanup time.Duration) *serverChainCache {
	return &serverChainCache{cache: cacheimpl.New(ttl, cleanup)}
}

func (s *serverChainCache) get(ip net.IP, port int, sni string, profile string) (string, bool) {
	key := fmt.Sprintf("%v_%v_%s_%s", ip, port, sni, profile)
	item, ok := s.cache.Get(key)
	if ok {
		return item.(string), true
	}
	return "", false
}

func (s *serverChainCache) set(ip net.IP, port int, sni string, profile string, chain string) {
	key := fmt.Sprintf("%v_%v_%s_%s", ip, port, sni, profile)
	s.cache.Set(key, chain, 0)
}

type uploadCache struct {
	cache *cacheimpl.Cache
}

func newUploadCache(ttl, cleanup time.Duration) *uploadCache {
	return &uploadCache{cache: cacheimpl.New(ttl, cleanup)}
}

// returns digest and key computed for cache
func (u *uploadCache) get(certs []*x509.Certificate) (string, string, bool) {
	_, key := certverify.DigestCerts(sha1.New(), certs)
	chain, ok := u.cache.Get(key)
	if ok {
		return chain.(string), key, true
	}
	return "", key, false
}

func (u *uploadCache) set(key, chain string) {
	u.cache.Set(key, chain, 0)
}

type verifyCache struct {
	cache *cacheimpl.Cache
}

func newVerifyCache(ttl, cleanup time.Duration) *verifyCache {
	return &verifyCache{cache: cacheimpl.New(ttl, cleanup)}
}

func (v *verifyCache) get(chain string, dnsname string) (tlsutil.VerifyResponse, bool) {
	key := fmt.Sprintf("%s_%s", chain, dnsname)
	item, ok := v.cache.Get(key)
	if ok {
		return item.(tlsutil.VerifyResponse), true
	}
	return tlsutil.VerifyResponse{}, false
}

func (v *verifyCache) set(chain string, dnsname string, response tlsutil.VerifyResponse) {
	key := fmt.Sprintf("%s_%s", chain, dnsname)
	v.cache.Set(key, response, 0)
}

type downloadCache struct {
	cache *cacheimpl.Cache
}

func newDownloadCache(ttl, cleanup time.Duration) *downloadCache {
	return &downloadCache{cache: cacheimpl.New(ttl, cleanup)}
}

func (d *downloadCache) get(chain string) ([]*x509.Certificate, bool) {
	item, ok := d.cache.Get(chain)
	if ok {
		return item.([]*x509.Certificate), true
	}
	return nil, false
}

func (d *downloadCache) set(chain string, certs []*x509.Certificate) {
	d.cache.Set(chain, certs, 0)
}
