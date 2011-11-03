package main

import (
	"testing"
	"log"
)

func checkDistance(t *testing.T, a, b Point, expected int) {
	distance := a.CrowDistance2(b)
	if distance != expected {
		t.Errorf("expected distance from %v to %v to be %v, got %v",
			a, b, expected, distance)
	}
}

func TestFilterPoints(t *testing.T) {

	testPoints := []Point{
		{1, 2}, {3, 4}, {4, 3}, {4, 4},
	}

	filtered := filterPoints(testPoints, func(p Point) bool {
		return p.c != 4
	})

	checkSamePoints(t, filtered, []Point{{1, 2}, {4, 3}})

	filtered = filterPoints(testPoints, func(p Point) bool {
		return p.r == 4 || p.c == 4
	})
	checkSamePoints(t, filtered, []Point{{3, 4}, {4, 3}, {4, 4}})

	// original unchanged
	checkSamePoints(t, testPoints, []Point{{1, 2}, {3, 4}, {4, 3}, {4, 4}})
}

func TestApplyToPoints(t *testing.T) {
	testPoints := []Point{
		{1, 2}, {3, 4}, {4, 3}, {4, 4},
	}
	applyToPoints(testPoints, func(p *Point) {
		p.r = p.c * 2
	})
	checkSamePoints(t, testPoints, []Point{{4, 2}, {8, 4}, {6, 3}, {8, 4}})
}

func checkSamePoints(t *testing.T, a, b []Point) {
	if len(a) != len(b) {
		t.Errorf("Expected %v and %v to be the same, but different lengths", a, b)
		return
	}
	for j, ap := range a {
		if !ap.Equals(b[j]) {
			t.Errorf("Expected %v and %v to be the same, but %v != %v", a, b, ap, b[j])
			return
		}
	}
}

func TestDistance(t *testing.T) {
	ROWS = 3
	COLS = 4
	o := Point{0, 0}
	data := []struct {
		a, b Point
		d    int
	}{
		{o, Point{2, 2}, 5},
		{o, Point{2, 3}, 2},
		{o, Point{0, 0}, 0},
		{o, Point{0, 1}, 1},
		{o, Point{1, 1}, 2},
		{o, Point{1, 0}, 1},
	}
	for _, d := range data {
		checkDistance(t, d.a, d.b, d.d)
		checkDistance(t, d.b, d.a, d.d)
	}
}

func TestNeighbours(t *testing.T) {
	ROWS = 40
	COLS = 30
	src := Point{10, 10}
	radius := 3
	n := src.Neighbours(radius)
	log.Printf("%v.Neighbours(%v) ->%v\n", src, radius, n)
}
