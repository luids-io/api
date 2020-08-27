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

// Buffer implements a buffer for async event notification.
type Buffer struct {
	//logger used for errors
	logger yalogi.Logger
	//notify
	notifier event.Notifier
	//data channel
	eventCh chan event.Event
	//control
	closed bool
	close  chan struct{}
}

// Option encapsules options for buffer.
type Option func(*bufferOpts)

type bufferOpts struct {
	logger yalogi.Logger
}

var defaultBufferOpts = bufferOpts{logger: yalogi.LogNull}

// SetLogger option allows set a custom logger.
func SetLogger(l yalogi.Logger) Option {
	return func(o *bufferOpts) {
		if l != nil {
			o.logger = l
		}
	}
}

// New returns a new event buffer.
func New(n event.Notifier, size int, opt ...Option) *Buffer {
	opts := defaultBufferOpts
	for _, o := range opt {
		o(&opts)
	}
	b := &Buffer{
		logger:   opts.logger,
		notifier: n,
		eventCh:  make(chan event.Event, size),
		close:    make(chan struct{}),
	}
	go b.doProcess()
	return b
}

// PushEvent implements an asyncronous notification.
func (b *Buffer) PushEvent(e event.Event) error {
	if b.closed {
		return errors.New("notifybuffer: buffer is closed")
	}
	b.eventCh <- e
	return nil
}

func (b *Buffer) doProcess() {
	for e := range b.eventCh {
		reqid, err := b.notifier.NotifyEvent(context.Background(), e)
		if err != nil {
			b.logger.Warnf("notifybuffer: notifying event with code '%v': %v", e.Code, err)
		}
		b.logger.Debugf("notifybuffer: notified event reqid: '%s'", reqid)
	}
	close(b.close)
}

// Close buffer
func (b *Buffer) Close() {
	if b.closed {
		return
	}
	b.closed = true
	close(b.eventCh)
	<-b.close
}
