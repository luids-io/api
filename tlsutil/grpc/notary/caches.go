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

func newServerChainCache(ttl int, cleanup time.Duration) *serverChainCache {
	return &serverChainCache{cache: cacheimpl.New(time.Duration(ttl)*time.Second, cleanup)}
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

func newUploadCache(ttl int, cleanup time.Duration) *uploadCache {
	return &uploadCache{cache: cacheimpl.New(time.Duration(ttl)*time.Second, cleanup)}
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
	ttl, negativettl int
	cache            *cacheimpl.Cache
}

func newVerifyCache(ttl, negativettl int, cleanup time.Duration) *verifyCache {
	return &verifyCache{
		ttl:         ttl,
		negativettl: negativettl,
		cache:       cacheimpl.New(time.Duration(ttl)*time.Second, cleanup),
	}
}

func (v *verifyCache) get(chain string, dnsname string) (tlsutil.VerifyResponse, bool) {
	key := fmt.Sprintf("%s_%s", chain, dnsname)
	hit, exp, ok := v.cache.GetWithExpiration(key)
	if ok {
		r := hit.(tlsutil.VerifyResponse)
		if r.TTL >= 0 {
			//updates ttl
			ttl := exp.Sub(time.Now()).Seconds()
			if ttl < 0 { //nonsense
				panic("cache missfunction")
			}
			r.TTL = int(ttl)
		}
		return r, true
	}
	return tlsutil.VerifyResponse{}, false
}

func (v *verifyCache) set(chain string, dnsname string, r tlsutil.VerifyResponse) tlsutil.VerifyResponse {
	//if don't cache
	if (r.TTL == tlsutil.NeverCache) || (!r.Invalid && v.negativettl == tlsutil.NeverCache) {
		return r
	}
	//sets cache
	ttl := v.ttl
	if !r.Invalid && v.negativettl > 0 {
		ttl = v.negativettl
	}
	if r.TTL > ttl { //if major than cachettl
		r.TTL = ttl //truncates reponse to cachettl
	}
	if r.TTL > 0 {
		key := fmt.Sprintf("%s_%s", chain, dnsname)
		v.cache.Set(key, r, time.Duration(r.TTL)*time.Second)
	}
	return r
}

type downloadCache struct {
	cache *cacheimpl.Cache
}

func newDownloadCache(ttl int, cleanup time.Duration) *downloadCache {
	return &downloadCache{cache: cacheimpl.New(time.Duration(ttl)*time.Second, cleanup)}
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
