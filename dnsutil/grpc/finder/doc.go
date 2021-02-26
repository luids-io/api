// Copyright 2021 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

// Package finder implements a dnsutil.Finder client and a ready to use
// service component.
//
// This package is a work in progress and makes no API stability promises.
package finder

import "fmt"

// Constants for api description.
const (
	APIName    = "luids.dnsutil"
	APIVersion = "v1"
	APIService = "Finder"
)

// ServiceName returns service name.
func ServiceName() string {
	return fmt.Sprintf("%s.%s.%s", APIName, APIVersion, APIService)
}
