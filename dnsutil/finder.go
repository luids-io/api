// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package dnsutil

import (
	"context"
	"net"
	"time"
)

// Finder is the interface for archive finder DNS information.
type Finder interface {
	GetResolv(ctx context.Context, id string) (ResolvData, bool, error)
	ListResolvs(ctx context.Context, filters []ResolvsFilter, rev bool, max int, next string) ([]ResolvData, string, error)
}

// ResolvsFilter stores filter information
type ResolvsFilter struct {
	Since, To      time.Time
	Server, Client net.IP
	Name           string
	ResolvedIP     net.IP
	ResolvedCNAME  string
	QID            int
	ReturnCode     int
	TLD            string
	TLDPlusOne     string
}
