package main

import (
	"fmt"
	"log"
)

// Try to move all the ants
func (m *Map) moveAll() {
	occupied := make(map[Location]bool)
	toMove := make([]*Ant, 0, len(m.myAnts))
	nextMove := make([]*Ant, 0, len(toMove)/4)

	for _, a := range m.myAnts {
		occupied[a.p.loc()] = true
		toMove = append(toMove, a)
		// log.Printf("Moving %s\n", a)
	}

	iterations, nMoved, deadlocked, blocked := 0, 0, 0, 0

	for len(toMove) > 0 {
		iterations += 1
		for _, a := range toMove {
			wantsMove, dst := a.WantsMove()
			if !wantsMove {
				continue
			}

			if m.isBlocked(dst) {
				log.Printf("%v blocked", dst)
				a.AbortMove()
				blocked += 1
				continue
			}
			if occupied[dst.loc()] {
				if dst.Equals(a.p) {
					a.Pause()
				} else {
					nextMove = append(nextMove, a)
				}
				continue
			}
			src := a.p
			a.Move(dst)
			occupied[dst.loc()] = true
			occupied[src.loc()] = false
			m.Moved(a, src, dst)
			nMoved += 1
		}

		// If we couldn't move any ants at all, we're deadlocked
		if len(toMove) == len(nextMove) {
			// deadlock
			log.Printf("deadlocked moves: %v", toMove)
			for _, a := range toMove {
				a.AbortMove()
				deadlocked += 1
			}
			nextMove = toMove[:0]
		}
		toMove, nextMove = nextMove[:], toMove[:0]
	}
	report := fmt.Sprintf("Moved %d in %d iterations", nMoved, iterations)
	if deadlocked > 0 {
		report += fmt.Sprintf(" %d deadlocked", deadlocked)
	}
	if blocked > 0 {
		report += fmt.Sprintf(" %d blocked", blocked)
	}
	log.Println(report)
}

func direction(src, dst Point) string {
	if !((src.r == dst.r) || (src.c == dst.c)) {
		log.Panicf("Cannot move from %v to %v\n", src, dst)
	}
	if dst.r == src.r+1 || (dst.r == 0 && src.r == ROWS-1) {
		return "s"
	}
	if dst.r == src.r-1 || (src.r == 0 && dst.r == ROWS-1) {
		return "n"
	}
	if dst.c == src.c+1 || (dst.c == 0 && src.c == COLS-1) {
		return "e"
	}
	if dst.c == src.c-1 || (src.c == 0 && dst.c == COLS-1) {
		return "w"
	}
	log.Panicf("Cannot move from %v to %v\n", src, dst)
	return ""
}

// We weren't able to move.  Give up
func (a *Ant) AbortMove() {
	a.plan = a.plan[:0]
}

// Does a want to move? where to?
func (a *Ant) WantsMove() (bool, Point) {
	if len(a.plan) == 0 {
		return false, a.p
	}
	return true, a.plan[0]
}

// If we can, make our move (and report success, update occupied)
func (a *Ant) Move(dst Point) {
	assert(dst.Equals(a.plan[0]), "dst: %v, a.plan: %v", dst, a.plan)

	fmt.Println("o", a.p.r, a.p.c, direction(a.p, dst))

	a.plan = a.plan[1:]
	a.p = dst
}

func (a *Ant) Pause() {
	assert (a.plan[0].Equals(a.p), "a.Pause: %v == %v", a.plan[0], a.p)
	a.plan = a.plan[1:]
}

func (m *Map) Moved(a *Ant, src, dst Point) {
	assert(m.myAnts[src.loc()] == a, "%v, %v", m.myAnts, src)
	assert(m.myAnts[dst.loc()] == nil, "%v, %v", m.myAnts, dst)

	m.myAnts[src.loc()] = nil, false
	m.myAnts[dst.loc()] = a
}

func assert(assertion bool, fmt string, fmtArgs ...interface{}) {
	if !assertion {
		log.Printf(fmt, fmtArgs...)
	}
}

// Assign a to get to p
// Return true if we have re-assigned an ant to get to 'p',
// false if we couldn't get ant there (or it was already en route)
func (a *Ant) moveTo(m *Map, p Point, reason string) bool {
	a.isTasked = true
	a.reason = reason

	// Do we already know how to get to p?
	if len(a.plan) > 0 && a.plan[len(a.plan)-1].Equals(p) {
		return false
	}

	path, error := m.ShortestPath(a.p, p)
	if error != nil {
		log.Printf("%v cannot get to %v (%s)\n", a, p, error)
		a.isTasked = false
		return false
	}
	a.plan = path
	return true
}

func (a *Ant) moveToPoint(m *Map, p Point, reason string) {
	a.isTasked = true
	a.reason = reason
	a.plan = []Point{p}
}
