// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

// Package classify implements a tlsutil.Classifier client and a ready to use
// service component.
//
// This package is a work in progress and makes no API stability promises.
package classify

import "fmt"

// Constants for api description.
const (
	APIName    = "luids.tlsutil"
	APIVersion = "v1"
	APIService = "Classify"
)

// ServiceName returns service name.
func ServiceName() string {
	return fmt.Sprintf("%s.%s.%s", APIName, APIVersion, APIService)
}
