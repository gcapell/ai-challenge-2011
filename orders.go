package main

import (
	"fmt"
	"log"
)

// Try to move all the ants
func (m *Map) moveAll() {
	iterations, nMoved, totalMoved := 0, 0, 0

	for  {
		iterations += 1
		nMoved = 0
		for _, a := range m.myAnts {
			if a.hasMoved || len(a.plan) == 0{
				continue
			}
			dst := a.plan[0]

			if m.isBlocked(dst)  {
				a.AbortMove(m, "blocked/food")
				continue
			}
			other := m.myAnts[dst.loc()]
			if other != nil {
				if other.hasMoved {
					a.AbortMove(m, "blocked/ant")
				}
				// Maybe this square will clear up later ...?
				continue
			}

			a.plan = a.plan[1:]
			a.Move(m, dst, a.reason)
			nMoved += 1
		}
		totalMoved += nMoved

		// If we couldn't move any ants at all, we're deadlocked (or finished)
		if nMoved == 0 {
			// deadlock
			for _, a := range m.myAnts {
				if !(a.hasMoved || len(a.plan)==0) {
					a.AbortMove(m, "deadlock")
				}
			}
			break
		}
	}
	log.Println(fmt.Sprintf("Moved %d in %d iterations, %v", totalMoved, iterations, m.movesThisTurn))
}

// We weren't able to move.  Give up
func (a *Ant) AbortMove(m *Map, reason string) {
	a.CancelPlans()
	a.Move(m, a.Point, reason)
}

func assert(assertion bool, fmt string, fmtArgs ...interface{}) {
	if !assertion {
		log.Panicf(fmt, fmtArgs...)
	}
}

// Assign a to get to p
// Return true if a is now on path to p
func (a *Ant) SetPathFor(m *Map, p Point, reason string) bool {
	a.reason = reason

	// Are we already en-route?
	if len(a.plan) > 0 && a.plan[len(a.plan)-1].Equals(p) {
		return true
	}

	path, error := a.ShortestPath(p, m)
	if error != nil {
		log.Printf("%v cannot get to %v (%s)\n", a, p, error)
		return false
	}
	a.plan = path
	return true
}

// Simultaneously move ants from 'src' to 'dst',
// overriding any other plans they may have had.
func SimultaneousOverridingMove(src, dst []Point, m *Map, reason string) {
	assert(noneMoved(m, src), "one of %v already moved", src)
	ants := make([]*Ant, len(src))
	for i, srcP := range src {
		dstP := dst[i]
		a := m.myAnts[srcP.loc()]
		ants[i] = a

		a.CancelPlans()
		a.MarkMoved()
		if !a.Equals(dstP) {
			a.OutputMove(dstP)
			a.Point = dstP
		}
		m.RecordMove(dstP, reason)
	}
	for _, p := range src {
		m.myAnts[p.loc()] = nil, false
	}
	for i, p := range dst {
		m.myAnts[p.loc()] = ants[i]
	}
}

func (a *Ant) MarkMoved() {
	assert (!a.hasMoved, "%v already moved", a)
	a.hasMoved = true
}

func (a *Ant) OverridingMove(m *Map, p Point, reason string) {
	a.CancelPlans()
	a.Move(m,p,reason)
}

func (a *Ant) Move(m *Map, p Point, reason string) {
	a.MarkMoved()
	if !a.Equals(p) {
		a.OutputMove(p)
		m.myAnts[a.loc()] = nil, false
		a.Point = p
		m.myAnts[p.loc()] = a
	}
	m.RecordMove(p, reason)
}

func (m *Map) RecordMove(dst Point, reason string) {
	points := m.movesThisTurn[reason]
	points = append(points, dst)
	m.movesThisTurn[reason] = points
}
