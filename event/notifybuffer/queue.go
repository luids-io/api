// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package notifybuffer

import (
	"errors"

	"github.com/luids-io/api/event"
)

// queue implements a buffer using a channel as a buffer
type queue struct {
	child   event.NotifyBuffer
	eventCh chan event.Event
	closed  bool
	waitc   chan struct{}
}

// NewQueue returns a new queue event buffer.
func NewQueue(child event.NotifyBuffer, size int) event.NotifyBuffer {
	q := &queue{
		child:   child,
		eventCh: make(chan event.Event, size),
		waitc:   make(chan struct{}),
	}
	go q.doProcess()
	return q
}

func (q *queue) doProcess() {
	for e := range q.eventCh {
		q.child.PushEvent(e)
	}
	close(q.waitc)
}

// PushEvent implements event.NotifyBuffer
func (q *queue) PushEvent(e event.Event) error {
	if q.closed {
		return errors.New("notifybuffer: buffer is closed")
	}
	q.eventCh <- e
	return nil
}

// Close implements event.NotifyBuffer
func (q *queue) Close() {
	if q.closed {
		return
	}
	q.closed = true
	close(q.eventCh)
	<-q.waitc
	//close child buffer
	q.child.Close()
}
