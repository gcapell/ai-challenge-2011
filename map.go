package main

import (
	"log"
	"fmt"
	"os"
	"strings"
)

type (
	//Item represents all the various items that may be on the map
	Item int8

	Square struct {
		isWater bool
		wasSeen bool	// Have we ever seen this square?
		lastSeen Turn	// .. if so, when?
	}

	Ant struct {
		p	Point		// Where are we now?
		plan	Points	// Where will we be?
		seen	Turn
		isBusy	bool
	}
	
	Map struct {
		squares	[][]Square
		
		myAnts	map[Location]*Ant
		myHills Points
		enemyHills Points
		
		enemies Points
		food	Points
		items	map[Location]Item
	}
)

const (
	UNKNOWN Item = iota - 5
	WATER
	FOOD
	LAND
	DEAD
	MY_ANT //= 0

	MAXPLAYER = 24
)

var (
	ROWS, COLS, VIEWRADIUS2 int
	TURN Turn
)

func (m *Map) Init(rows, cols, viewRadius2 int) {
	ROWS = rows
	COLS = cols
	VIEWRADIUS2 = viewRadius2

	m.myAnts = make(map[Location]*Ant)
	m.myHills = make([]Point, 0)
	m.enemies = make([]Point, 0)
	m.enemyHills = make([]Point, 0)
	m.food = make([]Point, 0)

	m.squares = make([][]Square, rows)
	for row:=0; row<rows; row++ {
		m.squares[row] = make([]Square, cols)
	}
	m.Reset()
}

//Reset clears the map for the next turn
func (m *Map) Reset() {
	m.myHills = m.myHills[:0]
	m.enemies = m.enemies[:0]
	m.enemyHills = m.enemyHills[:0]
	m.food = m.food[:0]
	m.items = make(map[Location]Item)
}

func (m *Map) isWet(p Point) bool {
	return m.squares[p.r][p.c].isWater
}

func (m *Map) DryNeighbours(p Point) []Point {
	allNeighbours := []Point {
		Point{p.r +1, p.c},
		Point{p.r -1, p.c},
		Point{p.r , p.c + 1},
		Point{p.r , p.c -1},
	}
	reply := make([]Point, 0, 4)
	for _, n := range(allNeighbours) {
		n.sanitise()
		if ! m.isWet(n) {
			reply = append(reply, n)
		}
	}
	return reply
}

// A scout has reported from 'p'
func (m *Map) ViewFrom(scout Point) {
	for _, p := range(scout.Neighbours(VIEWRADIUS2)) {
		s := &m.squares[p.r][p.c]
		s.wasSeen = true
		s.lastSeen = TURN
	}
}

func (m *Map) MarkWater(p Point) {
	m.squares[p.r][p.c].isWater = true
}

func (m *Map) MarkFood(p Point) {
	(&m.food).add(p)
	m.items[p.loc()] = FOOD
}

func (m *Map) MarkHill(p Point, ant Item) {
	if ant == MY_ANT {
		(&m.myHills).add(p)
	} else {
		(&m.enemyHills).add(p)
	}
}

func (m *Map) Update(words []string) {
	if words[0] == "turn" {
		turn  := Turn(atoi(words[1]))
		if turn != TURN+1 {
			log.Panicf("Turn number out of sync, expected %v got %v", TURN+1, turn)
		}
		TURN = turn
		return
	}

	p := Point{atoi(words[1]), atoi(words[2])}
	var ant Item
	if len(words) == 4 {
		ant = Item(atoi(words[3]))
	}

	switch words[0] {
	case "w":
		m.MarkWater(p)
	case "f":
		m.MarkFood(p)
	case "h":
		m.MarkHill(p, ant)
	case "a":
		m.AddAnt(p, ant)
	case "d":
		m.DeadAnt(p, ant)
	default:
		log.Panicf("unknown command updating map: %v\n", words)
	}
}

func (m *Map) AddAnt(p Point, ant Item) {
	if ant == MY_ANT {
		m.ViewFrom(p)
		
		antp, found := m.myAnts[p.loc()]
		if found { // existing ant?
			antp.seen = TURN
		} else { // new ant?
			m.myAnts[p.loc()] = &Ant{p:p, seen:TURN}
		}
	} else {
		m.enemies = append(m.enemies, p)
	}
	m.items[p.loc()] = ant
}

func (m *Map) DeadAnt(p Point, ant Item) {
	if ant != MY_ANT {
		return
	}
	m.myAnts[p.loc()] = nil, false
}

// We have received all the updates for this turn,
// make sure they make sense.
func (m *Map) UpdatesProcessed() {
	// Any ants that missed an update?
	for loc, ant := range(m.myAnts) {
		if (ant.seen != TURN) {
			log.Panicf("%v (@ %v) missed an update\n", ant, loc.point())
		}
	}
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
		for col, letter := range(line) {
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

func (m *Map) FreeAnts() []*Ant {
	reply := make([]*Ant, len(m.myAnts))
	for _, ant := range m.myAnts {
		if !ant.isBusy {
			reply = append(reply, ant)
		}
	}
	return reply
}

