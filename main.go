package main

import (
	"os"
	"log"
)

//main initializes the state and starts the processing loop
func main() {
	var s Game
	s.Start()
	mb := NewBot(&s)
	err := s.Loop(mb)
	if err != nil && err != os.EOF {
		log.Panicf("Loop() failed (%s)", err)
	}
}
