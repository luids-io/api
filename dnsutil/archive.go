// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package dnsutil

import (
	"context"
	"net"
	"time"

	"github.com/google/uuid"
)

// Archiver is the interface for archive DNS information.
type Archiver interface {
	SaveResolv(context.Context, ResolvData) (uuid.UUID, error)
}

// ResolvData stores information about DNS domain name resolutions.
type ResolvData struct {
	ID        uuid.UUID     `json:"id"`
	Timestamp time.Time     `json:"timestamp"`
	Duration  time.Duration `json:"duration"`
	Server    net.IP        `json:"server"`
	Client    net.IP        `json:"client"`
	//query
	QID        uint16           `json:"qid"`
	Name       string           `json:"name"`
	IsIPv6     bool             `json:"isIPv6"`
	QueryFlags ResolvQueryFlags `json:"queryFlags"`
	//response
	ReturnCode     int                 `json:"returnCode"`
	ResolvedIPs    []net.IP            `json:"resolvedIPs,omitempty"`
	ResolvedCNAMEs []string            `json:"resolvedCNAMEs,omitempty"`
	ResponseFlags  ResolvResponseFlags `json:"responseFlags"`
	//tld calculated info
	TLD        string `json:"tld"`
	TLDPlusOne string `json:"tldPlusOne"`
}

// ResolvQueryFlags stores information about resolv query flags
type ResolvQueryFlags struct {
	Do                bool `json:"do"`
	AuthenticatedData bool `json:"authenticatedData"`
	CheckingDisabled  bool `json:"checkingDisabled"`
}

// ResolvResponseFlags stores information about resolv response flags
type ResolvResponseFlags struct {
	AuthenticatedData bool `json:"authenticatedData"`
}
