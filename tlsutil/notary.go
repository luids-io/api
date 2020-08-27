// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package tlsutil

import (
	"context"
	"crypto/x509"
	"net"
)

// Notary is the main interface that must be implemented by notary services.
type Notary interface {
	GetServerChain(ctx context.Context, ip net.IP, port int, sni, profile string) (string, error)
	VerifyChain(ctx context.Context, chain string, dnsname string, force bool) (VerifyResponse, error)
	DownloadCerts(ctx context.Context, chain string) ([]*x509.Certificate, error)
	//danger methods
	SetServerChain(ctx context.Context, ip net.IP, port int, sni, profile string, chain string) error
	UploadCerts(ctx context.Context, certs []*x509.Certificate) (string, error)
}

//VerifyResponse stores information about the service's verification responses.
type VerifyResponse struct {
	// Invalid is true if the chain and dnsname is invalid
	Invalid bool `json:"invalid"`
	// Reason stores the reason why it's invalid
	Reason string `json:"reason,omitempty"`
	// TTL is a number in seconds used for caching
	TTL int `json:"ttl"`
}

// NeverCache is a special value for TTL. If TTLs has this value, caches
// should not store the response.
const NeverCache = -1
