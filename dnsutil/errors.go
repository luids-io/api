// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package dnsutil

import "errors"

// Some standard errors returned by interfaces.
var (
	ErrCanceledRequest = errors.New("dnsutil: canceled request")
	ErrBadRequest      = errors.New("dnsutil: bad request")
	ErrNotSupported    = errors.New("dnsutil: not supported")
	ErrUnavailable     = errors.New("dnsutil: not available")
	ErrInternal        = errors.New("dnsutil: internal error")
	//limit errors
	ErrLimitDNSClientQueries = errors.New("dnsutil: max queries per dns client")
	ErrLimitResolvedNamesIP  = errors.New("dnsutil: max names resolved for an ip")
)
