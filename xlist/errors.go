// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package xlist

import "errors"

// Some standard errors returned by List interfaces
var (
	ErrCanceledRequest = errors.New("xlist: canceled request")
	ErrBadRequest      = errors.New("xlist: bad request")
	ErrNotSupported    = errors.New("xlist: resource not supported")
	ErrUnavailable     = errors.New("xlist: not available")
	ErrInternal        = errors.New("xlist: internal error")
)
