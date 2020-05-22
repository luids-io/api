// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package tlsutil

import "errors"

// Some standard errors returned by interfaces
var (
	ErrCanceledRequest = errors.New("canceled request")
	ErrBadRequest      = errors.New("bad request")
	ErrNotSupported    = errors.New("not supported")
	ErrUnavailable     = errors.New("not available")
	ErrInternal        = errors.New("internal error")
	ErrTimeOutOfSync   = errors.New("time out of sync")
	// specific interface Analyze
	ErrStreamNotFound   = errors.New("stream not found")
	ErrDuplicatedStream = errors.New("duplicated stream")
)
