package main

import (
	"testing"
	"strings"
	"log"
	"os"
	"fmt"
)

func (m *Map) InitFromString(viewRadius2 int, s string) os.Error {
	lines := strings.Fields(s)
	rows := len(lines)
	var cols int
	for row, line := range lines {
		if row == 0 {
			cols = len(line)
			m.Init(rows, cols, viewRadius2)
		} else {
			if cols != len(line) {
				return fmt.Errorf("different-length lines in %v", lines)
			}
		}
		for col, letter := range line {
			p := Point{row, col}
			switch letter {
			case '#':
				// Unknown territory
			case '%':
				m.MarkWater(p)
			case '*':
				m.MarkFood(p)
			case 'a':
				m.AddAnt(p, 0)
			case 'b':
				m.AddAnt(p, 1)
			}
		}
	}

	return nil
}

func loadMap() Map {
	var m Map
	m.Init(4, 3, 0)
	return m
}

func TestInitFromString(t *testing.T) {
	var m Map
	s := `
		###
		#%#
		a#b
	`
	m.InitFromString(0, s)
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
	m.InitFromString(0, `
		..%..%..
		a.%.b%..
		..%.....
		..%.....
	`)
	src := Point{1, 0} // a
	dst := Point{1, 4} // b
	path, err := m.ShortestPath(src, dst)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%v -> %v : %v\n", src, dst, path)
}

func TestLocationConversion(t *testing.T) {
	loc := Point{3, 2}.loc()
	p := loc.point()
	if p.r != 3 || p.c != 2 {
		t.Errorf("conversion broken, got %v, wanted (3, 2)", p)
	}
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
