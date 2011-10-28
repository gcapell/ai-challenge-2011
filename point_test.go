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
