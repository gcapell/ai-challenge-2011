package main

import (
	"math"
)

type (
	//Location combines (Row, Col) coordinate pairs 
	// for use as keys in maps (and in a 1d array)
	Location int

	Point  struct{ r, c int } // rows, columns
	Points []Point
)

func (s *Points) add(p Point) {
	*s = append(*s, p)
}

func (loc Location) point() Point {
	iLoc := int(loc)
	return Point{iLoc / COLS, iLoc % COLS}
}

func sanitisePoint(p *Point) {
	p.sanitise()
}

func (p *Point) sanitise() {
	if p.r < 0 {
		p.r += ROWS
	}
	if p.r >= ROWS {
		p.r -= ROWS
	}
	if p.c < 0 {
		p.c += COLS
	}
	if p.c >= COLS {
		p.c -= COLS
	}
}

func (p Point) In(other []Point) bool {
	for _, o := range other {
		if p.Equals(o) {
			return true
		}
	}
	return false
}

func (p Point) sanitised() Point {
	p.sanitise()
	return p
}

func (p Point) loc() Location {
	return Location(p.r*COLS + p.c)
}

func (p Point) Equals(r Point) bool {
	return p.r == r.r && p.c == r.c
}

func wrapDelta(a, b, wrap int) int {
	delta := abs(a - b)
	return min(delta, wrap-delta)
}

// Return (Manhattan) distance between two points,
// allowing for warping across edges
func (a Point) Distance(b Point) (int, int) {
	return wrapDelta(a.r, b.r, ROWS), wrapDelta(a.c, b.c, COLS)
}

func (a Point) CrowDistance2(b Point) int {
	dx, dy := a.Distance(b)
	return dx*dx + dy*dy
}

func (a Point) CrowDistance(b Point) float64 {
	return math.Sqrt(float64(a.CrowDistance2(b)))
}

// Given slice of points, and predicate function,
// return slice of points where predicate is true.
func filterPoints(points []Point, fn func(Point) bool) []Point {
	success := 0

	reply := make([]Point, len(points))
	for j, p := range points {
		if fn(p) {
			reply[success] = points[j]
			success += 1
		}
	}
	return reply[:success]
}

// Apply (modifying) function to points in slice
func applyToPoints(points []Point, fn func(*Point)) {
	for j := range points {
		fn(&points[j])
	}
}

func (p Point) NeighboursAndSelf() []Point {
	reply := []Point{
		p,
		{p.r + 1, p.c},
		{p.r - 1, p.c},
		{p.r, p.c + 1},
		{p.r, p.c - 1},
	}
	applyToPoints(reply, sanitisePoint)
	return reply
}

func (p Point) Neighbours(rad2 int) []Point {
	reply := Points(make([]Point, rad2))
	if rad2 < 1 {
		return reply
	}
	for dr := 0; dr*dr <= rad2; dr++ {
		for dc := 0; dc*dc+dr*dr <= rad2; dc++ {
			reply.add(Point{p.r + dr, p.c + dc}.sanitised())
			if dr != 0 {
				reply.add(Point{p.r - dr, p.c + dc}.sanitised())
			}
			if dc != 0 {
				reply.add(Point{p.r + dr, p.c - dc}.sanitised())
			}
			if dr != 0 && dc != 0 {
				reply.add(Point{p.r - dr, p.c - dc}.sanitised())
			}
		}
	}
	return reply
}

