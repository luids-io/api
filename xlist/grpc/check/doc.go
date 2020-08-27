// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

// Package check implements a xlist.Checker client and a ready to use service
// component.
//
// This package is a work in progress and makes no API stability promises.
package check

import "fmt"

// Constants for api description
const (
	APIName    = "luids.xlist"
	APIVersion = "v1"
	APIService = "Check"
)

// ServiceName returns service name
func ServiceName() string {
	return fmt.Sprintf("%s.%s.%s", APIName, APIVersion, APIService)
}
