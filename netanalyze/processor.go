// Copyright 2020 Luis Guillén Civera <luisguillenc@gmail.com>. All rights reserved.

package netanalyze

import (
	"github.com/google/gopacket"
)

// Processor defines main Process function
type Processor interface {
	Process(name string, source gopacket.PacketDataSource, hooks *Hooks) (stop func(), errs <-chan error, err error)
}
