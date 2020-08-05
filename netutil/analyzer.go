// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package netutil

import (
	"fmt"
	"time"
)

// AnalyzerFactory interface is used for create Analyzer services
type AnalyzerFactory interface {
	NewAnalyzer(name string) (Analyzer, error)
}

// Analyzer interface defines analyzer methods
type Analyzer interface {
	SendPacket(layer Layer, md PacketMetadata, data []byte) error
	Close() error
}

// Layer of packets
type Layer int

// Layer types
const (
	Ethernet Layer = iota
	IPv4
	IPv6
)

func (l Layer) String() string {
	switch l {
	case Ethernet:
		return "eth"
	case IPv4:
		return "ip4"
	case IPv6:
		return "ip6"
	default:
		return fmt.Sprintf("unkown(%d)", l)
	}
}

// PacketMetadata is a copy of gopacket.PacketMetadata
type PacketMetadata struct {
	// Timestamp is the time the packet was captured, if that is known.
	Timestamp time.Time
	// CaptureLength is the total number of bytes read off of the wire.
	CaptureLength int
	// Length is the size of the original packet.  Should always be >= CaptureLength.
	Length int
	// InterfaceIndex
	InterfaceIndex int
}
