// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

// Package notify implements a event.Notifier client and a ready to use
// service component.
//
// This package is a work in progress and makes no API stability promises.
package notify

import "fmt"

// Constants for api description.
const (
	APIName    = "luids.event"
	APIVersion = "v1"
	APIService = "Notify"
)

// ServiceName returns service name.
func ServiceName() string {
	return fmt.Sprintf("%s.%s.%s", APIName, APIVersion, APIService)
}
