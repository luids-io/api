// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package tlsutil

import "errors"

// Some standard errors returned by interfaces.
var (
	ErrCanceledRequest = errors.New("tlsutil: canceled request")
	ErrBadRequest      = errors.New("tlsutil: bad request")
	ErrNotSupported    = errors.New("tlsutil: not supported")
	ErrUnavailable     = errors.New("tlsutil: not available")
	ErrInternal        = errors.New("tlsutil: internal error")
	// specific interface Analyze
	ErrTimeOutOfSync    = errors.New("tlsutil: time out of sync")
	ErrMsgOutOfOrder    = errors.New("tlsutil: message time out of order")
	ErrStreamNotFound   = errors.New("tlsutil: stream not found")
	ErrDuplicatedStream = errors.New("tlsutil: duplicated stream")
	ErrAnalyzerExists   = errors.New("tlsutil: analyzer with name exists")
	// specific interface Notary
	ErrDialingWithServer = errors.New("tlsutil: dialing with server")
	ErrChainNotFound     = errors.New("tlsutil: chain not found")
	ErrCertNotFound      = errors.New("tlsutil: certificate not found")
)
