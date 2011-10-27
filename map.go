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

	Move struct {
		src Location
		d   Direction
	}

	//Location combines (Row, Col) coordinate pairs for use as keys in maps (and in a 1d array)
	Location int

	//Direction represents the direction concept for issuing orders.
	Direction int

	Point struct {r, c int}	// rows, columns

	Square struct {
		isWater bool
		wasSeen bool	// Have we ever seen this square?
		lastSeen Turn	// .. if so, when?
	}

	Ant struct {
		p	Point
		plan	[] Point
		target	Item	// food, enemy ant, enemy hill, ...
	}
	
	Map struct {
		squares	[][]Square
		
		myAnts	map[Location]Ant
		myHills []Point
		
		enemies []Point
		enemyHills []Point
		food	[]Point
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

const (
	North Direction = iota
	East
	South
	West

	NoMovement
)

var (
	ROWS, COLS, VIEWRADIUS2 int
	TURN Turn
	DIRS = []Direction{North, East, South, West}
	directionstrings = map[Direction]string{
		North:      "n",
		South:      "s",
		West:       "w",
		East:       "e",
		NoMovement: "-",
	}
)

func (loc Location ) point () Point {
	iLoc := int(loc)
	return Point{ iLoc / COLS, iLoc % COLS}
}

func (p *Point) sanitise() {
	if p.r < 0 {
		p.r += ROWS
	}
	if p.r >= ROWS {
		p.r -= ROWS
	}
	if p.c < 0 {
		p.c += COLS
	}
	if p.c >= COLS {
		p.c -= COLS
	}
}

func (p Point) loc() Location {
	p.sanitise()
	return Location(p.r * COLS + p.c)
}

func (p Point) Equals(r Point) bool {
	return p.r == r.r && p.c == r.c
}

func (m *Map) Init(rows, cols, viewRadius2 int) {
	ROWS = rows
	COLS = cols
	VIEWRADIUS2 = viewRadius2

	m.myAnts = make(map[Location]Ant)
	m.myHills = make([]Point)
	n.enemies = make([]Point)
	m.enemyHills = make([]Point)
	m.food = make([]Point)


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
	for _, p := range(m.Neighbours(scout, VIEWRADIUS2)) {
		s := &m.squares[p.r][p.c]
		s.wasSeen = true
		s.lastSeen = TURN
	}
}

type PointSlice []Point

func (s *PointSlice) add(p Point) {
	*s = append(*s, p)
}

func (m *Map) Neighbours(p Point, rad2 int) [] Point{
	reply := PointSlice(make([]Point, rad2))
	if rad2 < 1 {
		return reply
	}
	for dr:=0; dr*dr<= rad2; dr++ {
		for dc := 0; dc*dc + dr  *dr <= rad2; dc++ {
			reply.add(Point{p.r + dr, p.c + dc})
			if(dr != 0) {
				reply.add(Point{p.r - dr, p.c + dc})
			}
			if (dc !=0) {
				reply.add(Point{p.r + dr, p.c - dc})
			}
			if (dr !=0 && dc != 0) {
				reply.add(Point{p.r - dr, p.c - dc})
			}
		}
	}
	return reply
}

func (d Direction) String() string {
	return directionstrings[d]
}

//Move returns a new location which is one step in the specified direction from the specified location.
func (m *Map) Move(p Point, d Direction) Point {
	switch d {
	case North:
		p.r -= 1
	case South:
		p.r += 1
	case West:
		p.c -= 1
	case East:
		p.c += 1
	case NoMovement: //do nothing
	default:
		log.Panicf("%v is not a valid direction", d)
	}
	p.sanitise()
	return p
}

func (m *Map) MarkWater(p Point) {
	m.squares[p.r][p.c].isWater = true
}

func (m *Map) MarkFood(p Point) {
	m.Food[p.loc()] = TURN
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
		m.Hills[p.loc()] = ant
	case "a":
		m.AddAnt(p, ant)
	case "d":
		// Ignore dead ant
	default:
		log.Panicf("unknown command: %v\n", words)
	}
}

func (m *Map) AddAnt(p Point, ant Item) {
	if ant == MY_ANT {
		loc = p.loc()
		m.ViewFrom(p)
		// existing ant?
		
		// new ant?
		
	} else {
		m.enemies = append(m.enemies, p)
	}
}


//Call IssueOrderLoc to issue an order for an ant at loc
func (m *Map) IssueOrderLoc(p Point, d Direction) {
	//...
	fmt.Println("o", p.r, p.c, d)
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

func wrapDelta(a, b, wrap int) int {
	delta := a-b
	if delta<0 {
		delta = -delta
	}
	wrapped := wrap - delta
	// log.Printf("a: %d, b: %d, wrap: %d, delta: %d, wrapped: %d\n", a, b, wrap, delta, wrapped)
	if delta < wrapped {
		return delta
	}
	return wrapped
}


// Return (Manhattan) distance between two points,
// allowing for warping across edges
func (m *Map) Distance( a,b Point) int {
	return wrapDelta(a.r, b.r, ROWS)+
		wrapDelta(a.c, b.c, COLS)
}
