package main

import (
	"testing"
	"strings"
	"fmt"
	"os"
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

func (m *Map) InitFromString(s string, viewRadius2 int) os.Error {
	lines := strings.Fields(s)
	rows := len(lines)
	var cols int
	for row, line := range(lines) {
		if row == 0 {
			cols = len(line)
			m.Init(rows, cols, viewRadius2)
		} else {
			if cols != len(line) {
				return fmt.Errorf("different-length lines in %v" , lines)
			}
		}
		fmt.Printf("line:[%s]\n", line)
		for col, letter := range(line) {
			loc := m.FromRowCol(row, col)
			switch letter {
			case '%':
				m.MarkWater(loc)
			case '*':
				m.MarkFood(loc)
			case 'a':
				m.AddAnt(loc, 0)
			case 'b':
				m.AddAnt(loc, 1)
			}
		}
	}

	return nil
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
