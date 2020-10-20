// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. See LICENSE.

package notifybuffer_test

import (
	"sync"
	"testing"
	"time"

	"github.com/luids-io/api/event"
	"github.com/luids-io/api/event/notifybuffer"
)

func TestWaitDups(t *testing.T) {
	notifier := &mockNotifier{}

	buffer := notifybuffer.NewWaitDups(notifier, 10, 100*time.Millisecond)

	e1 := event.New(1234, event.Low)
	e1.Set("prueba", "valor1")
	e2 := event.New(1234, event.Low)
	e2.Set("prueba", "valor2")
	e3 := event.New(1235, event.Low)
	e3.Set("prueba", "valor1")

	buffer.PushEvent(e1)
	buffer.PushEvent(e2)
	buffer.PushEvent(e1)
	buffer.PushEvent(e1)
	oevents := notifier.GetEvents()
	if len(oevents) > 0 {
		t.Fatalf("events before wait: %v", oevents)
	}
	// waiting for exced time
	time.Sleep(120 * time.Millisecond)
	oevents = notifier.GetEvents()
	if len(oevents) != 2 {
		t.Fatalf("events missmatch: %v", oevents)
	}
	if e1.Code != oevents[0].Code || e1.Level != oevents[0].Level || e1.PrintFields() != oevents[0].PrintFields() {
		t.Fatalf("unexpected event[0]: %v", oevents[0])
		if oevents[0].Duplicates != 2 {
			t.Fatalf("duplicates of event[0] missmatch: %v", oevents[0].Duplicates)
		}
	}
	if e2.Code != oevents[1].Code || e2.Level != oevents[1].Level || e2.PrintFields() != oevents[1].PrintFields() {
		t.Fatalf("unexpected event[1]: %v", oevents[1])
		if oevents[1].Duplicates != 0 {
			t.Fatalf("duplicates of event[1] missmatch: %v", oevents[1].Duplicates)
		}
	}
	buffer.PushEvent(e3)
	buffer.PushEvent(e3)
	buffer.PushEvent(e2)
	buffer.PushEvent(e2)
	// close buffer forcing push events
	buffer.Close()

	oevents = notifier.GetEvents()
	if len(oevents) != 2 {
		t.Fatalf("events missmatch: %v", oevents)
	}
	if e3.Code != oevents[0].Code || e3.Level != oevents[0].Level || e3.PrintFields() != oevents[0].PrintFields() {
		t.Fatalf("unexpected event[0]: %v", oevents[0])
		if oevents[0].Duplicates != 1 {
			t.Fatalf("duplicates of event[0] missmatch: %v", oevents[0].Duplicates)
		}
	}
	if e2.Code != oevents[1].Code || e2.Level != oevents[1].Level || e2.PrintFields() != oevents[1].PrintFields() {
		t.Fatalf("unexpected event[1]: %v", oevents[1])
		if oevents[1].Duplicates != 1 {
			t.Fatalf("duplicates of event[1] missmatch: %v", oevents[1].Duplicates)
		}
	}

}

func TestWaitDupsFilled(t *testing.T) {
	notifier := &mockNotifier{}

	buffer := notifybuffer.NewWaitDups(notifier, 2, 100*time.Millisecond)

	e1 := event.New(1234, event.Low)
	e1.Set("prueba", "valor1")
	e2 := event.New(1234, event.Low)
	e2.Set("prueba", "valor2")
	e3 := event.New(1235, event.Low)
	e3.Set("prueba", "valor1")

	buffer.PushEvent(e1)
	buffer.PushEvent(e1)
	buffer.PushEvent(e2)
	buffer.PushEvent(e3)
	oevents := notifier.GetEvents()
	if len(oevents) != 1 {
		t.Fatalf("unexecpected events before wait: %v", oevents)
	}
	if oevents[0].Code != 1235 {
		t.Fatalf("unexpected output event: %v", oevents[0])
	}
	// waiting for exced time
	time.Sleep(120 * time.Millisecond)
	oevents = notifier.GetEvents()
	if len(oevents) != 2 {
		t.Fatalf("events missmatch: %v", oevents)
	}
	if e1.Code != oevents[0].Code || e1.Level != oevents[0].Level || e1.PrintFields() != oevents[0].PrintFields() {
		t.Fatalf("unexpected event[0]: %v", oevents[0])
		if oevents[0].Duplicates != 1 {
			t.Fatalf("duplicates of event[0] missmatch: %v", oevents[0].Duplicates)
		}
	}
	if e2.Code != oevents[1].Code || e2.Level != oevents[1].Level || e2.PrintFields() != oevents[1].PrintFields() {
		t.Fatalf("unexpected event[1]: %v", oevents[1])
		if oevents[1].Duplicates != 0 {
			t.Fatalf("duplicates of event[1] missmatch: %v", oevents[1].Duplicates)
		}
	}
	// close buffer forcing push events
	buffer.Close()
	// buffer must be empty
	oevents = notifier.GetEvents()
	if len(oevents) != 0 {
		t.Fatalf("events missmatch: %v", oevents)
	}
}

func TestWaitDupsFragmented(t *testing.T) {
	notifier := &mockNotifier{}

	buffer := notifybuffer.NewWaitDups(notifier, 3, 100*time.Millisecond)

	e1 := event.New(1111, event.Low)
	e2 := event.New(2222, event.Low)
	e3 := event.New(3333, event.Low)
	e4 := event.New(4444, event.Low)

	buffer.PushEvent(e1)
	buffer.PushEvent(e1)
	buffer.PushEvent(e1)
	// waiting for a fragment time
	time.Sleep(80 * time.Millisecond)
	buffer.PushEvent(e2)
	time.Sleep(40 * time.Millisecond)
	buffer.PushEvent(e2)
	buffer.PushEvent(e3)
	buffer.PushEvent(e4)
	buffer.PushEvent(e3)
	oevents := notifier.GetEvents()
	if len(oevents) != 1 {
		t.Fatalf("events before wait: %v", oevents)
	}
	if oevents[0].Code != e1.Code || oevents[0].Duplicates != 2 {
		t.Fatalf("unexpected event[0]: %v %v", oevents[0].Code, oevents[0].Duplicates)
	}
	// exceed buffer
	buffer.PushEvent(e1)
	oevents = notifier.GetEvents()
	if len(oevents) != 1 {
		t.Fatalf("events before wait: %v", oevents)
	}
	if oevents[0].Code != e1.Code || oevents[0].Duplicates != 0 {
		t.Fatalf("unexpected event[0]: %v", oevents[0])
	}
	// closes buffer
	buffer.Close()
	oevents = notifier.GetEvents()
	if len(oevents) != 3 {
		t.Fatalf("events after close: %v", oevents)
	}
	if oevents[0].Code != e3.Code || oevents[0].Duplicates != 1 {
		t.Fatalf("unexpected event[0]: %v", oevents[0].Code)
	}
	if oevents[1].Code != e2.Code || oevents[1].Duplicates != 1 {
		t.Fatalf("unexpected event[1]: %v", oevents[1].Code)
	}
	if oevents[2].Code != e4.Code || oevents[2].Duplicates != 0 {
		t.Fatalf("unexpected event[2]: %v", oevents[2].Code)
	}
}

// mockups
type mockNotifier struct {
	mu     sync.Mutex
	events []event.Event
}

func (m *mockNotifier) PushEvent(e event.Event) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.events = append(m.events, e)
	return nil
}

func (m *mockNotifier) GetEvents() []event.Event {
	m.mu.Lock()
	defer m.mu.Unlock()
	ret := make([]event.Event, len(m.events), len(m.events))
	copy(ret, m.events)
	m.events = m.events[:0]
	return ret
}

func (m *mockNotifier) Close() {}
