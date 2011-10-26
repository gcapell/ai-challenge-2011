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
	loc := toLoc(3, 2)
	row, col := toRC(loc)
	if row != 3 || col != 2 {
		t.Errorf("conversion broken, got (%v, %v), wanted (3, 2)", row, col)
	}
}

func TestMove(t *testing.T) {
	m := loadMap()

	loc := toLoc(3, 2)

	n := toLoc(2, 2)
	s := toLoc(0, 2)
	e := toLoc(3, 0)
	w := toLoc(3, 1)

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
	src := toLoc(row, col)
	n := m.Neighbours( src, radius)
	fmt.Printf("(%d,%d),r:%d ->", row, col, radius)
	for _, loc := range(n) {
		row, col = toRC(loc)
		fmt.Printf("(%d,%d), ", row, col)
	}
	fmt.Printf("\n")
}

func TestMap(t *testing.T) {
	m := loadMap()

	m.AddAnt(toLoc(2, 2), MY_ANT)
	m.AddAnt(toLoc(0, 2), MY_ANT)

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
