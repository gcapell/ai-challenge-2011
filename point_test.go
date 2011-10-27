package main
import (
	"testing"
	"fmt"
)

func TestDistance(t *testing.T) {
	ROWS=3
	COLS=4
	o := Point{0,0}
	data := []struct {a,b Point; d int} {
		{o, Point{2,2}, 3},
		{o, Point{2,3}, 2},
		{o, Point{0,0}, 0},
		{o, Point{0,1}, 1},
		{o, Point{1,1}, 2},
		{o, Point{1,0}, 1},
	}
	for _, d := range(data) {
		distance := d.a.Distance( d.b)
		if distance != d.d {
			t.Errorf("expected distance from %v to %v to be %d, got %d",
				d.a, d.b, d.d, distance)
		}
		distance = d.b.Distance(d.a)
		if distance != d.d {
			t.Errorf("expected distance from %v to %v to be %d, got %d",
				d.b, d.a, d.d, distance)
		}
	}
}

func TestNeighbours(t *testing.T) {
	ROWS=40
	COLS=30
	src := Point{10,10}
	radius := 3
	n := src.Neighbours( radius)
	fmt.Printf("%v.Neighbours(%v) ->%v\n", src, radius, n)
}

