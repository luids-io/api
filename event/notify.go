// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package event

import (
	"context"
	"errors"
	"os"
	"path/filepath"
)

// Notifier is the interface for event notifiers.
type Notifier interface {
	NotifyEvent(ctx context.Context, e Event) (string, error)
}

// NotifyBuffer is the interface for event buffer implementations.
type NotifyBuffer interface {
	PushEvent(e Event) error
}

// SetBuffer sets the default buffer instance.
func SetBuffer(b NotifyBuffer) {
	defaultBuffer = b
}

// Notify notifies using the default buffer instance.
func Notify(e Event) error {
	if defaultBuffer != nil {
		return defaultBuffer.PushEvent(e)
	}
	return errors.New("event: default buffer not defined")
}

// SetDefaultInstance changes default instance name.
func SetDefaultInstance(label string) {
	if instanceRegExp.MatchString(label) {
		defaultSource.Instance = label
	}
}

// GetDefaultSource returns default notify events source.
func GetDefaultSource() Source {
	return defaultSource
}

var defaultSource Source
var defaultBuffer NotifyBuffer

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	defaultSource = Source{
		Hostname: hostname,
		Program:  filepath.Base(os.Args[0]),
		PID:      os.Getpid(),
	}
}
