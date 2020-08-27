// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package netutil

import "errors"

// Some standard errors returned by interfaces.
var (
	ErrBadRequest   = errors.New("netutil: bad request")
	ErrNotSupported = errors.New("netutil: not supported")
	ErrUnavailable  = errors.New("netutil: not available")
	ErrInternal     = errors.New("netutil: internal error")
	// specific interface Analyze
	ErrAnalyzerExists   = errors.New("netutil: analyzer with name exists")
	ErrTimeOutOfSync    = errors.New("netutil: time out of sync")
	ErrPacketOutOfOrder = errors.New("netutil: packet time out of order")
)
