// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

// Package forward implements a event.Forwarder client and a ready to use
// service component.
//
// This package is a work in progress and makes no API stability promises.
package forward

import "fmt"

// Constants for api description.
const (
	APIName    = "luids.event"
	APIVersion = "v1"
	APIService = "Forward"
)

// ServiceName returns service name.
func ServiceName() string {
	return fmt.Sprintf("%s.%s.%s", APIName, APIVersion, APIService)
}
