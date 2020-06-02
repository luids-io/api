// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package tlsutil

import (
	"context"
	"crypto/x509"
	"net"
	"time"
)

// Notary is the main interface that must be implemented by notary services
type Notary interface {
	GetServerChain(ctx context.Context, ip net.IP, port int, sni, profile string) (string, error)
	VerifyChain(ctx context.Context, chain string, dnsname string, force bool) (VerifyResponse, error)
	DownloadCerts(ctx context.Context, chain string) ([]*x509.Certificate, error)
	//danger methods
	SetServerChain(ctx context.Context, ip net.IP, port int, sni, profile string, chain string) error
	UploadCerts(ctx context.Context, certs []*x509.Certificate) (string, error)
}

//VerifyResponse stores information about the service's verification responses
type VerifyResponse struct {
	// Timestamp stores verify time response
	Timestamp time.Time `json:"timestamp"`
	// Invalid is true if the chain and dnsname is invalid
	Invalid bool `json:"invalid"`
	// Reason stores the reason why it's not valid
	Reason string `json:"reason,omitempty"`
}
