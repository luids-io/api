// Copyright 2019 Luis Guillén Civera <luisguillenc@gmail.com>. View LICENSE.

package check

import (
	"testing"
	"time"

	"github.com/luids-io/api/xlist"
)

func TestCacheWithoutBounds(t *testing.T) {
	c := newCache(0, 0, 1*time.Minute)
	var tests = []struct {
		name  string
		r     xlist.Response
		sleep time.Duration
		want  bool
		wttl  int
	}{
		{"1.1.1.1", xlist.Response{TTL: 0}, 0, false, 0},
		{"1.1.1.1", xlist.Response{TTL: 1}, 0, true, 1},
		{"1.1.1.1", xlist.Response{TTL: 1}, 2 * time.Second, false, 0},
		{"1.1.1.1", xlist.Response{TTL: 2}, 1 * time.Second, true, 1},
	}
	for _, test := range tests {
		c.set(test.name, xlist.IPv4, test.r)
		if test.sleep > 0 {
			time.Sleep(test.sleep)
		}
		r, got := c.get(test.name, xlist.IPv4)
		if test.want != got || test.wttl != r.TTL {
			t.Errorf("test %v fails, got %v %v", test, got, r)
		}
	}
}

func TestCacheWithBounds(t *testing.T) {
	c := newCache(1, 2, 1*time.Minute)
	var tests = []struct {
		name  string
		r     xlist.Response
		sleep time.Duration
		want  bool
		wttl  int
	}{
		{"1.1.1.1", xlist.Response{TTL: 0}, 0, true, 0},
		{"1.1.1.1", xlist.Response{TTL: 1}, 0, true, 1},
		{"1.1.1.1", xlist.Response{TTL: 1}, 2 * time.Second, false, 0},
		{"1.1.1.1", xlist.Response{TTL: 2}, 1 * time.Second, true, 1},
		{"1.1.1.1", xlist.Response{TTL: 4}, 3 * time.Second, false, 0},
	}
	for _, test := range tests {
		c.set(test.name, xlist.IPv4, test.r)
		if test.sleep > 0 {
			time.Sleep(test.sleep)
		}
		r, got := c.get(test.name, xlist.IPv4)
		if test.want != got || test.wttl != r.TTL {
			t.Errorf("test %v fails, got %v %v", test, got, r)
		}
	}
}
