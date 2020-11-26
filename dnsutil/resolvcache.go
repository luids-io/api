// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package dnsutil

import (
	"context"
	"net"
	"time"
)

// ResolvCache interface defines a cache for dns resolutions.
type ResolvCache interface {
	ResolvCollector
	ResolvChecker
}

// ResolvCollector interface collects to the cache resolved ip address
// information.
type ResolvCollector interface {
	Collect(ctx context.Context, client net.IP, name string, resolved []net.IP, cnames []string) error
}

// ResolvChecker is the interface for checks in a resolv cache.
// Params client and resolved are required, name can be empty.
type ResolvChecker interface {
	Check(ctx context.Context, client, resolved net.IP, name string) (CacheResponse, error)
}

// CacheResponse stores cache response information.
type CacheResponse struct {
	// Result is true if was resolved
	Result bool `json:"result"`
	// Last time resolved
	Last time.Time `json:"last,omitempty"`
	// Store time of cache
	Store time.Time `json:"store"`
}
