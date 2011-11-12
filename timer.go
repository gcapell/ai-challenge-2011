package main

import (
	"log"
	"os"
)

type (
	Timer interface {
		Reset()
		Split(s string)
	}
	
	ConcreteTimer struct {
		start int64
		split int64
	}
	NullTimer int
)

var NULLTIMER *NullTimer

func NewTimer() Timer {
	return new(ConcreteTimer)
}

func (t *ConcreteTimer) Reset() {
	now := nsec()
	t.start = now
	t.split = now
}

func (t *ConcreteTimer) Split(s string) {
	now := nsec()

	deltaSplit := float64(now-t.split) / 1e9
	delta := float64(now-t.start) / 1e9
	t.split = now

	if true {
		log.Printf("%s: %.3f %.3f", s, deltaSplit, delta)
	}
}

func nsec() int64 {
	s, ns, _ := os.Time()
	return s*1e9 + ns
}

func NewNullTimer() Timer {
	return NULLTIMER
}
func (n *NullTimer) Reset() {}
func (n *NullTimer) Split(s string) {}

