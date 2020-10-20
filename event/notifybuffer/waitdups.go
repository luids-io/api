// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package notifybuffer

import (
	"errors"
	"sync"
	"time"

	"github.com/luids-io/api/event"
)

type waitdups struct {
	child    event.NotifyBuffer
	mu       sync.Mutex
	duration time.Duration
	waiting  int
	items    []witem
	closed   bool
	close    chan struct{}
	waitc    chan struct{}
}

type witem struct {
	waiting   bool
	timestamp time.Time
	event     event.Event
	fields    string
}

// NewWaitDups returns a new event buffer.
func NewWaitDups(child event.NotifyBuffer, size int, d time.Duration) event.NotifyBuffer {
	w := &waitdups{
		child:    child,
		duration: d,
		items:    make([]witem, size, size),
		close:    make(chan struct{}),
		waitc:    make(chan struct{}),
	}
	go w.doSync()
	return w
}

func (w *waitdups) doSync() {
	tick := time.NewTicker(w.duration)
	defer tick.Stop()
SYNCLOOP:
	for {
		select {
		case <-tick.C:
			w.sendExpired()
		case <-w.close:
			break SYNCLOOP
		}
	}
	w.sendAll()
	close(w.waitc)
}

// PushEvent implements event.NotifyBuffer
func (w *waitdups) PushEvent(e event.Event) error {
	if w.closed {
		return errors.New("waitdups: buffer is closed")
	}
	w.waitEvent(e)
	return nil
}

// Close implements event.NotifyBuffer
func (w *waitdups) Close() {
	if w.closed {
		return
	}
	w.closed = true
	close(w.close)
	<-w.waitc
	//close child buffer
	w.child.Close()
}

func (w *waitdups) waitEvent(e event.Event) {
	fields := e.PrintFields()
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now()
	// item buffer is empty
	if w.waiting == 0 {
		w.items[0].waiting = true
		w.items[0].timestamp = now
		w.items[0].event = e
		w.items[0].fields = fields
		w.waiting = 1
		return
	}
	// events waiting
	freeidx := -1
	seen := 0
SEARCH_LOOP:
	for idx, item := range w.items {
		if !item.waiting {
			if freeidx < 0 {
				freeidx = idx
			}
			continue
		}
		//item not waiting, compare
		if item.event.Code == e.Code && item.event.Level == e.Level && item.fields == fields {
			w.items[idx].event.Duplicates++
			return
		}
		seen++
		if seen >= w.waiting {
			//all waiting items has been seen
			if freeidx < 0 {
			ALLOC_NEXT_FREE:
				for nidx := idx + 1; nidx < len(w.items); nidx++ {
					if !w.items[nidx].waiting {
						freeidx = nidx
						break ALLOC_NEXT_FREE
					}
				}
				//not found free index
				if freeidx < 0 {
					break SEARCH_LOOP
				}
			}
			//free item available
			w.items[freeidx].waiting = true
			w.items[freeidx].timestamp = now
			w.items[freeidx].event = e
			w.items[freeidx].fields = fields
			w.waiting++
			return
		}
	}
	//not available items
	w.child.PushEvent(e)
}

func (w *waitdups) sendAll() {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.waiting == 0 {
		return
	}
	for idx, item := range w.items {
		if item.waiting {
			w.child.PushEvent(item.event)
			w.items[idx].waiting = false
			w.waiting--
		}
	}
}

func (w *waitdups) sendExpired() {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.waiting == 0 {
		return
	}
	now := time.Now()
	for idx, item := range w.items {
		if item.waiting && now.Sub(item.timestamp) > w.duration {
			w.child.PushEvent(item.event)
			w.items[idx].waiting = false
			w.waiting--
		}
	}
}
