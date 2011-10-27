package main

import (
	"fmt"
	"log"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West

	NoMovement
)

var (
	DIRS = []Direction{North, East, South, West}
	directionstrings = map[Direction]string{
		North:      "n",
		South:      "s",
		West:       "w",
		East:       "e",
		NoMovement: "-",
	}
)

func (d Direction) String() string {
	return directionstrings[d]
}

//Call IssueOrderLoc to issue an order for an ant at loc
func (m *Map) IssueOrderLoc(p Point, d Direction) {
	//...
	fmt.Println("o", p.r, p.c, d)
}

// Assign a to get to p
func (a *Ant) moveTo(m *Map, p Point) {
	a.isBusy = true
	
	// Do we already know how to get to p?
	if len(a.plan) > 0 && a.plan[len(a.plan)-1].Equals(p) {
		log.Printf("%v using cached plan %v\n", a, a.plan)
		return
	}

	path, error := m.ShortestPath(a.p, p)
	if error != nil {
		log.Printf("%v cannot get to %v (%s)\n", a, p, error)
		a.isBusy = false
	}
	a.plan = path
}
