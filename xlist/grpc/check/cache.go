// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package check

import (
	"fmt"
	"time"

	cacheimpl "github.com/patrickmn/go-cache"

	"github.com/luids-io/api/xlist"
)

const defaultCacheCleanups = 3 * time.Minute

// cache implements a client cache
// client cache stores responses using defined ttls
// minttl and maxttl are bounds to time storage in cache
type cache struct {
	minTTL int
	maxTTL int
	cachei *cacheimpl.Cache
}

type cacheitem struct {
	stored time.Time
	r      xlist.Response
}

// newCache returns a cache
func newCache(minTTL, maxTTL int, cleanups time.Duration) *cache {
	c := &cache{
		minTTL: minTTL,
		maxTTL: maxTTL,
		cachei: cacheimpl.New(cacheimpl.NoExpiration, cleanups),
	}
	return c
}

// Flush cleas all items from cache
func (c *cache) flush() {
	c.cachei.Flush()
}

func (c *cache) get(name string, resource xlist.Resource) (xlist.Response, bool) {
	key := fmt.Sprintf("%s_%s", resource.String(), name)
	hit, ok := c.cachei.Get(key)
	if ok {
		item := hit.(cacheitem)
		if item.r.TTL >= 0 {
			now := time.Now()
			//updates ttl with time stored in cache
			fttl := now.Sub(item.stored).Seconds()
			if fttl < 0 { //nonsense
				panic("cache missfunction")
			}
			ttl := item.r.TTL - int(fttl)
			if ttl >= 0 {
				item.r.TTL = ttl
			} else {
				item.r.TTL = 0
			}
		}
		return item.r, true
	}
	return xlist.Response{}, false
}

func (c *cache) set(name string, resource xlist.Resource, r xlist.Response) xlist.Response {
	//if don't cache
	if r.TTL == xlist.NeverCache {
		return r
	}
	//sets ttl in bounds
	ttl := r.TTL
	if c.minTTL > 0 && ttl < c.minTTL {
		ttl = c.minTTL
	}
	if c.maxTTL > 0 && ttl > c.maxTTL {
		ttl = c.maxTTL
	}
	//client cache doesn't change reponse ttl, it only caches
	if ttl > 0 {
		key := fmt.Sprintf("%s_%s", resource.String(), name)
		c.cachei.Set(key, cacheitem{stored: time.Now(), r: r}, time.Duration(ttl)*time.Second)
	}
	return r
}
