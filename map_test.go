package main

import (
	"testing"
	"strings"
	"fmt"
)

func loadMap() Map {
	var m Map
	m.Init(4,3,0)
	return m
}

func TestDistance(t *testing.T) {
	m := loadMap()
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
		distance := m.Distance(d.a, d.b)
		if distance != d.d {
			t.Errorf("expected distance from %v to %v to be %d, got %d",
				d.a, d.b, d.d, distance)
		}
		distance = m.Distance(d.b, d.a)
		if distance != d.d {
			t.Errorf("expected distance from %v to %v to be %d, got %d",
				d.b, d.a, d.d, distance)
		}
	}
}

func TestInitFromString(t *testing.T)  {
	var m Map 
	s := `
		###
		#%#
		a#b
	`
	m.InitFromString(s, 0)
	checkMap(t, &m, "InitFromString", s)
}

func TestPrint(t *testing.T) {
	m := loadMap()
	checkMap(t, &m, "map loads/prints wrong size", `
		###
		###
		###
		###
	`)
}

func TestShortestPath(t *testing.T) {
	var m Map
	m.InitFromString(`
		..%..%..
		a.%.b%..
		..%.....
		..%.....
	`, 0)
	src := Point{1,0}	// a
	dst := Point{1,4}	// b
	path, err := m.ShortestPath(src.loc(), dst.loc())
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v -> %v : %v\n", src, dst, path)
}

func TestLocationConversion(t *testing.T) {
	loc := Point{3, 2}.loc()
	p := loc.point()
	if p.r != 3 || p.c != 2{
		t.Errorf("conversion broken, got %v, wanted (3, 2)", p)
	}
}

func TestMove(t *testing.T) {
	m := loadMap()

	src := Point{3, 2}

	data := []struct{direction Direction; dst Point} {
		{North, Point{2,2}},
		{South, Point{0, 2}},
		{East, Point{3, 0}},
		{West, Point{3, 1}},
	}
	for _, d := range(data) {
		actualDst := m.Move(src, d.direction)
		if !actualDst.Equals(d.dst) {
			t.Errorf("Move %v expected %v, got %v", d.direction, d.dst, actualDst)
		}
	}
}

func TestNeighbours(t *testing.T) {
	var m Map
	m.Init(40,30,0)
	row, col, radius := 10,10,3
	src := Point{row, col}
	n := m.Neighbours( src, radius)
	fmt.Printf("(%d,%d),r:%d ->", row, col, radius)
	for _, p := range(n) {
		fmt.Printf("%v, ", p)
	}
	fmt.Printf("\n")
}

func TestMap(t *testing.T) {
	m := loadMap()

	m.AddAnt(Point{2, 2}, MY_ANT)
	m.AddAnt(Point{0, 2}, MY_ANT)

	checkMap(t, &m, "ants in wrong place", `
		##a
		###
		##a
		###
	`)
}

func checkMap(t *testing.T, m *Map, msg string, expected string) {
	if !sameText(m.String(), expected) {
		t.Errorf("%s, expected: %s, got:\n%s\n", msg, expected, m.String())
	}
}

// Are a and b the same (excluding whitespace)?
func sameText(a, b string) bool {
	return stripWhite(a) == stripWhite(b)
}

func stripWhite(s string) string {
	return strings.Join(strings.Fields(s), "\n")
}
