// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package xlist

import (
	"context"
)

// Checker is the interface for check lists.
type Checker interface {
	// Check method checks if the value encoded as string is in the list
	Check(ctx context.Context, name string, r Resource) (Response, error)
	// Resources returns an array with the resource types supported by the RBL service
	Resources(ctx context.Context) ([]Resource, error)
}

//Response stores information about the service's responses.
type Response struct {
	// Result is true if the value is in the list
	Result bool `json:"result"`
	// Reason stores the reason why it is the list (or not if you want)
	Reason string `json:"reason,omitempty"`
	// TTL is a number in seconds used for caching
	TTL int `json:"ttl"`
}

// NeverCache is a special value for TTL. If TTLs has this value, caches
// should not store the response.
const NeverCache = -1
