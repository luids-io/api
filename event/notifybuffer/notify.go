// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

// Package notifybuffer provides a simple event.NotifyBuffer implementation.
//
// This package is a work in progress and makes no API stability promises.
package notifybuffer

import (
	"context"
	"errors"

	"github.com/luids-io/api/event"
	"github.com/luids-io/core/yalogi"
)

type notifier struct {
	logger   yalogi.Logger
	notifier event.Notifier
	closed   bool
}

// Notifier wrappes a Notifier and implements an event.NotifyBuffer.
func Notifier(n event.Notifier, l yalogi.Logger) event.NotifyBuffer {
	return &notifier{logger: l, notifier: n}
}

// PushEvent implements event.NotifyBuffer
func (b *notifier) PushEvent(e event.Event) error {
	if b.closed {
		return errors.New("notifier is closed")
	}
	reqid, err := b.notifier.NotifyEvent(context.Background(), e)
	if err != nil {
		b.logger.Warnf("notifybuffer: notifying event with code '%v': %v", e.Code, err)
		return err
	}
	b.logger.Debugf("notifybuffer: notified event '%s'", reqid)
	return nil
}

// Close implements event.NotifyBuffer
func (b *notifier) Close() {
	if b.closed {
		return
	}
	b.closed = true
}
