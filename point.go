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

func countBool(slice []bool) int {
	count := 0
	for _, b := range slice {
		if b {
			count += 1
		}
	}
	return count
}

func (p Point) In (other []Point) bool  {
	for _, o := range other {
		if p.Equals(o) {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sign(a int) int {
	if a < 0 {
		return -1
	}
	return 1
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

// FIXME - eventually cope with warping
func between(a, b Point) Point {
	return Point{(a.r + b.r) / 2, (a.c + b.c) / 2}
}

// Is 'p' close enough to 'between' a and b?
func (p Point) intercepts(a, b Point) bool {
	if p.CrowDistance(a) < 4 || p.CrowDistance(b) < 4 {
		return true
	}
	rowRatio, rowBetween := linear(a.r, p.r, b.r)
	columnRatio, columnBetween := linear(a.c, p.c, b.c)

	return rowBetween && columnBetween && similarRatio(rowRatio, columnRatio)
}

func similarRatio(r1, r2 float64) bool {
	return r1 == 0 || r2 == 0 || math.Fabs(r1-r2) < .05
}

func linear(a, p, b int) (ratio float64, between bool) {
	if abs(a-b) < 4 {
		return 0, abs(p-(a+b)/2) < 2
	}
	if sign(b-a) != sign(b-p) {
		return 0, false
	}
	if sign(b-a) != sign(p-a) {
		return 0, false
	}

	return float64(abs(b-p)) / float64(abs(b-a)), true
}

func intercept(defender, enemy, victim Point) Point {
	if defender.intercepts(victim, enemy) {
		return enemy
	}
	return between(victim, enemy)
}
