// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package netutil

import "errors"

// Some standard errors returned by interfaces
var (
	ErrBadRequest   = errors.New("bad request")
	ErrNotSupported = errors.New("not supported")
	ErrUnavailable  = errors.New("not available")
	ErrInternal     = errors.New("internal error")
	// specific interface Analyze
	ErrTimeOutOfSync = errors.New("time out of sync")
)
