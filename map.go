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
		isTasked	bool	// Has this ant been given an order this turn?
		hasMoved	bool	// Has this ant moved already this turn?
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
func (a *Ant) Distance(p Point) int {
	return a.p.Distance(p)
}

func (m *Map) nearestHill(p Point) Point {
	var closest Point
	for i, hill := range m.myHills {
		if i==0 {
			closest = hill
		} else {
			if p.Distance(hill) < p.Distance(closest) {
				closest = hill
			}
		}
	}
	return closest
}


func (m*Map) EnemiesNearOurHill(tooClose int) []Point{
	reply := make([]Point, 0)

	enemyLoop: for _, e := range m.enemies {
		for _, hill := range m.myHills {
			if e.Distance(hill) <= tooClose {
				reply = append(reply, e)
				continue enemyLoop
			}
		}
	}
	return reply
}

func (a *Ant) String() string {
	if len(a.plan) > 0 {
		return fmt.Sprintf("Ant@%v->%v", a.p, a.plan[len(a.plan)-1])
	} 
	return fmt.Sprintf("Ant@%v", a.p)
}

const (
	UNKNOWN Item = iota - 5
	WATER
	FOOD
	LAND
	DEAD
	MY_ANT = 0

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
		ant.isTasked = false	// open for business!
		ant.hasMoved = false
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

// Return slice of ants who aren't already assigned
// If interrupt, include ants who already are moving
func (m *Map) FreeAnts(interrupt bool) []*Ant {
	reply := make([]*Ant, 0, len(m.myAnts))
	for _, ant := range m.myAnts {
		if ant.isTasked {
			// Already assigned this turn
			continue
		}
		if len(ant.plan) > 0 && !interrupt {
			// Ant already has a plan
			continue
		}
		reply = append(reply, ant)
	}
	return reply
}

