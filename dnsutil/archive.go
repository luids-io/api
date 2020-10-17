// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package dnsutil

import (
	"context"
	"net"
	"time"
)

// Archiver is the interface for archive DNS information.
type Archiver interface {
	SaveResolv(context.Context, *ResolvData) (string, error)
}

// ResolvData stores information about DNS domain name resolutions.
type ResolvData struct {
	ID        string        `json:"id"`
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
	ReturnCode    int                 `json:"returnCode"`
	ResolvedIPs   []net.IP            `json:"resolvedIPs,omitempty"`
	ResponseFlags ResolvResponseFlags `json:"responseFlags"`
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
