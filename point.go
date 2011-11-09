package main

import (
	"math"
	"log"
	"fmt"
)

type (
	//Location combines (Row, Col) coordinate pairs 
	// for use as keys in maps (and in a 1d array)
	Location uint

	Point  struct{ r, c int } // rows, columns
)

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

func (a Point) CrowDistance2(b Point) int {
	dx, dy := wrapDelta(a.r, b.r, ROWS), wrapDelta(a.c, b.c, COLS)
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

type PointQueue struct  {
	pos int
	points []Point
}

func NewPointQueue(n int) PointQueue {
	return PointQueue{0, make([]Point, n)}
}

func (pq *PointQueue) Add (points... Point) {
	for _, p := range points {
		pq.points[pq.pos] = p
		pq.pos += 1
	}
}

func (pq *PointQueue) SanitiseAndExport () []Point {
	exported := pq.points[:pq.pos]
	applyToPoints(exported, sanitisePoint)
	return exported
}

func (p Point) Neighbours(rad2 int) []Point {
	if rad2 < 1 {
		return make([]Point,0)
	}
	pq := NewPointQueue(4*rad2 +1)
	for dr := 0; dr*dr <= rad2; dr++ {
		for dc := 0; dc*dc+dr*dr <= rad2; dc++ {
			pq.Add(Point{p.r + dr, p.c + dc})
			if dr != 0 {
				pq.Add(Point{p.r - dr, p.c + dc})
			}
			if dc != 0 {
				pq.Add(Point{p.r + dr, p.c - dc})
			}
			if dr != 0 && dc != 0 {
				pq.Add(Point{p.r - dr, p.c - dc})
			}
		}
	}
	return pq.SanitiseAndExport()
}

func (p Point) spiral(step, maxDistance int) []Point {
	pointsAcross := (2*maxDistance/step) + 1
	pq := NewPointQueue(pointsAcross*pointsAcross)
	for radius := step; radius < maxDistance; radius += step {
		for off := 0; off < radius; off += step {
			pq.Add(
				Point{p.r + radius - off, p.c - radius},
				Point{p.r - radius, p.c - radius + off},
				Point{p.r - radius + off, p.c + radius},
				Point{p.r + radius, p.c + radius - step},
			)
		}
	}
	return pq.SanitiseAndExport()
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

func (src Point) OutputMove(dst Point) {
	fmt.Println("o", src.r, src.c, direction(src, dst))
}


func  minDistance2(src Point, other []Point ) int {
	var minD int
	for i, p := range other {
		d := src.CrowDistance2(p)
		if i== 0 || d<minD {
			minD = d
		}
	}
	return minD
}
