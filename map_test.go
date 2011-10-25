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

func TestLocationConversion(t *testing.T) {
	m := loadMap()
	loc := m.FromRowCol(3, 2)
	row, col := m.FromLocation(loc)
	if row != 3 || col != 2 {
		t.Errorf("conversion broken, got (%v, %v), wanted (3, 2)", row, col)
	}

	loc2 := m.FromRowCol(3, -1)
	if loc2 != loc {
		t.Errorf("from xy broken, got (%v), wanted (%v)", loc2, loc)
	}

}

func TestMove(t *testing.T) {
	m := loadMap()

	loc := m.FromRowCol(3, 2)

	n := m.FromRowCol(2, 2)
	s := m.FromRowCol(4, 2)
	e := m.FromRowCol(3, 3)
	w := m.FromRowCol(3, 1)

	if n != m.Move(loc, North) {
		t.Errorf("Move north is broken")
	}
	if s != m.Move(loc, South) {
		t.Errorf("Move south is broken")
	}
	if e != m.Move(loc, East) {
		t.Errorf("Move east is broken")
	}
	if w != m.Move(loc, West) {
		t.Errorf("Move west is broken")
	}
}

func TestNeighbours(t *testing.T) {
	var m Map
	m.Init(40,30,0)
	row, col, radius := 10,10,3
	src := m.FromRowCol(row, col)
	n := m.Neighbours( src, radius)
	fmt.Printf("(%d,%d),r:%d ->", row, col, radius)
	for _, loc := range(n) {
		row, col = m.FromLocation(loc)
		fmt.Printf("(%d,%d), ", row, col)
	}
	fmt.Printf("\n")
}

func TestMap(t *testing.T) {
	m := loadMap()

	m.AddAnt(m.FromRowCol(2, 2), MY_ANT)
	m.AddAnt(m.FromRowCol(4, 2), MY_ANT)

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
