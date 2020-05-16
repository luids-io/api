// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package tlsutil

import (
	"fmt"
	"net"
	"time"
)

// AnalyzerFactory interface is used for create Analyzer services
type AnalyzerFactory interface {
	NewAnalyzer(name string) (Analyzer, error)
}

// Analyzer interface defines analyzer methods
type Analyzer interface {
	SendMessage(m *Msg) error
	Close() error
}

// Msg defines message for analyzer
type Msg struct {
	Type     MsgType
	StreamID int64
	Open     *MsgOpen
	Data     *MsgData
}

func (m *Msg) String() string {
	s := fmt.Sprintf("type=%v,streamid=%v", m.Type, m.StreamID)
	if m.Open != nil {
		s = s + "," + m.Open.String()
	}
	if m.Data != nil {
		s = s + "," + m.Data.String()
	}
	return s
}

// MsgType defines message types
type MsgType int8

// Type possible values
const (
	DataMsg MsgType = iota
	OpenMsg
	CloseMsg
)

func (m MsgType) String() string {
	switch m {
	case DataMsg:
		return "data"
	case OpenMsg:
		return "open"
	case CloseMsg:
		return "close"
	default:
		return "unknown"
	}
}

// MsgOpen stores required data by the open message
type MsgOpen struct {
	SrcIP, DstIP     net.IP
	SrcPort, DstPort int
}

func (m *MsgOpen) String() string {
	return fmt.Sprintf("srcip=%v,srcport=%v,dstip=%v,dstport=%v",
		m.SrcIP, m.SrcPort, m.DstIP, m.DstPort)
}

// MsgData stores required data by the data message
type MsgData struct {
	Timestamp        time.Time
	Bytes            int
	SawStart, SawEnd bool
	Records          [][]byte
	Error            error
}

func (m *MsgData) String() string {
	return fmt.Sprintf("ts=%v,bytes=%v,sawstart=%v,sawend=%v,records=%v,error=%v",
		m.Timestamp, m.Bytes, m.SawStart, m.SawEnd, len(m.Records), m.Error)
}
