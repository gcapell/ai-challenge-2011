package main

import (
	"testing"
	"strings"
	"fmt"
	"os"
)

func loadMap() Map {
	var m Map
	var g Game
	g.Rows = 4
	g.Cols = 3
	m.Init(&g)
	m.Reset()
	return m
}

func TestInitFromString(t *testing.T)  {
	var m Map 
	m.InitFromString(`
		...
		.%.
		a.b
	`)
}

func (m *Map) InitFromString(s string) os.Error {
	lines := strings.Fields(s)
	rows := len(lines)
	var cols int
	for i, line := range(lines) {
		if i == 0 {
			cols = len(line)
		} else {
			if cols != len(line) {
				return fmt.Errorf("different-length lines in %v" , lines)
			}
		}
		fmt.Printf("line:[%s]\n", line)
	}
	m.Rows = rows
	m.Cols = cols
	return nil
}

func TestPrint(t *testing.T) {
	m := loadMap()
	if m.String() != `. . . 
. . . 
. . . 
. . . 
` {
		t.Errorf("map loads/prints wrong size, got `%s`", m)
	}
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

	if m.String() != `. . a 
. . . 
. . a 
. . . 
` {
		t.Errorf("map put ants in wrong place, got `%s`", m.String())
	}
}
