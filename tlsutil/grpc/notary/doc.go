// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

// Package notary provides a notary client and a ready-to-use service
// wrapper for the luids.tlsutil.vX.Notary API
//
// This package is a work in progress and makes no API stability promises.
package notary

import "fmt"

// Constants for api description
const (
	APIName    = "luids.tlsutil"
	APIVersion = "v1"
	APIService = "Notary"
)

// ServiceName returns service name
func ServiceName() string {
	return fmt.Sprintf("%s.%s.%s", APIName, APIVersion, APIService)
}
