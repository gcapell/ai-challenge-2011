package main

import (
	"fmt"
	"log"
)

// Try to move all the ants
func (m *Map) moveAll() {
	occupied := make(map[Location]bool)
	for _, a := range m.myAnts {
		occupied[a.p.loc()] = true
	}

	for {
		nMoved := 0
		for _, a := range m.myAnts {
			if a.Move(m, occupied) {
				nMoved += 1
			}
		}
		log.Printf("moveAll moved %d", nMoved)
		if nMoved == 0 {
			break
		}
	}
}

func direction(src, dst Point) string {
	if !((src.r == dst.r) || (src.c == dst.c)) {
		log.Panicf("Cannot move from %v to %v\n", src, dst)
	}
	if dst.r == src.r+1 || (src.r == 0 && dst.r == ROWS-1) {
		return "s"
	}
	if dst.r == src.r-1 || (dst.r == 0 && src.r == ROWS-1) {
		return "n"
	}
	if dst.c == src.c+1 || (src.c == 0 && dst.c == COLS-1) {
		return "e"
	}
	if dst.c == src.c-1 || (dst.c == 0 && src.c == COLS-1) {
		return "w"
	}
	log.Panicf("Cannot move from %v to %v\n", src, dst)
	return ""

}

// If we can, make our move (and report success, update occupied)
func (a *Ant) Move(m *Map, occupied map[Location]bool) bool {
	if a.hasMoved || len(a.plan) == 0 {
		return false
	}

	dst := a.plan[0]
	if occupied[dst.loc()] {
		return false
	}

	// Report the move
	src := a.p
	fmt.Println("o", src.r, src.c, direction(src, dst))

	// Update our status
	a.hasMoved = true
	a.plan = a.plan[1:]
	a.p = dst

	occupied[src.loc()] = false
	occupied[dst.loc()] = true

	m.myAnts[src.loc()] = nil, false
	m.myAnts[dst.loc()] = a

	return true
}

// Assign a to get to p
func (a *Ant) moveTo(m *Map, p Point, reason string) {
	log.Printf("Moving %v to %v for %s", a, p, reason)
	a.isTasked = true

	// Do we already know how to get to p?
	if len(a.plan) > 0 && a.plan[len(a.plan)-1].Equals(p) {
		log.Printf("%v using cached plan %v\n", a, a.plan)
		return
	}

	path, error := m.ShortestPath(a.p, p)
	if error != nil {
		log.Printf("%v cannot get to %v (%s)\n", a, p, error)
		a.isTasked = false
		return
	}
	a.plan = path
}
