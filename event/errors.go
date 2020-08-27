// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package event

import "errors"

// Some standard errors returned by interfaces.
var (
	ErrCanceledRequest = errors.New("event: canceled request")
	ErrBadRequest      = errors.New("event: bad request")
	ErrUnauthorized    = errors.New("event: unauthorized")
	ErrNotSupported    = errors.New("event: not supported")
	ErrUnavailable     = errors.New("event: not available")
	ErrInternal        = errors.New("event: internal error")
)
