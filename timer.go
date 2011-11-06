package main

import (
	"log"
)

type Timer struct {
	start int64
	split int64
}

func (t *Timer) Reset() {
	now := nsec()
	t.start = now
	t.split = now
}

func (t *Timer) Split(s string) {
	now := nsec()

	deltaSplit := float64(now-t.split) / 1e9
	delta := float64(now-t.start) / 1e9
	t.split = now

	if true {
		log.Printf("%s: %.3f %.3f", s, deltaSplit, delta)
	}
}

