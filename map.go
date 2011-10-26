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

	Map struct {
		squares	[][]Square

		Ants         map[Location]Item
		Hills        map[Location]Item
		Food         map[Location]Turn
		Destinations map[Location]bool
		MyAnts       map[Location]bool // ant location -> is moving?
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

	m.Food = make(map[Location]Turn)
	m.Hills = make(map[Location]Item)

	m.squares = make([][]Square, rows)
	for row:=0; row<rows; row++ {
		m.squares[row] = make([]Square, cols)
	}
	m.Reset()
}

//Reset clears the map for the next turn
func (m *Map) Reset() {
	m.Ants = make(map[Location]Item)
	m.Destinations = make(map[Location]bool)
	m.MyAnts = make(map[Location]bool)
}

// Given start location, return slice of legal next points
func (m *Map) NextValidMoves(p Point) []Point {
	next := make([]Point, 0, 4)

	for _, d := range DIRS {
		p2 := m.Move(p, d)
		if m.SafeDestination(p2) {
			next = append(next, p2)
		}
	}
	return next
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
		if ! m.squares[n.r][n.c].isWater {
			reply = append(reply, n)
		}
	}
	return reply
}


func (m *Map) EnemyHillAt(p Point) bool {
	item, found := m.Hills[p.loc()]
	return found && item != MY_ANT
}

func (m *Map) FoodAt(p Point) bool {
	_, found := m.Food[p.loc()]
	return found
}

func (m *Map) MyStationaryAnts() chan Location {
	ch := make(chan Location)
	go func() {
		for loc, isMoving := range m.MyAnts {
			if !isMoving {
				ch <- loc
			}
		}
		close(ch)
	}()
	return ch
}

//ViewFrom adds a circle of land centered on the given location
func (m *Map) ViewFrom(center Point) {
	for _, p := range(m.Neighbours(center, VIEWRADIUS2)) {
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

func (m *Map) AddDestination(p Point) {
	if m.Destinations[p.loc()] {
		log.Panicf("Already have something at that destination!")
	}
	m.Destinations[p.loc()] = true
}

func (m *Map) RemoveDestination(p Point) {
	m.Destinations[p.loc()] = false, false
}

//SafeDestination will tell you if the given location is a 
//safe place to dispatch an ant. It considers water and both
//ants that have already sent an order and those that have not.
func (m *Map) SafeDestination(p Point) bool {
	return !m.squares[p.r][p.c].isWater && !m.Destinations[p.loc()]
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
	m.Ants[p.loc()] = ant

	//if it turns out that you don't actually use the visible radius for anything,
	//feel free to comment this out. It's needed for the image debugging, though.
	if ant == MY_ANT {
		m.AddDestination(p)
		m.ViewFrom(p)
		m.MyAnts[p.loc()] = false
	}
}


//Call IssueOrderLoc to issue an order for an ant at loc
func (m *Map) IssueOrderLoc(p Point, d Direction) {
	dest := m.Move(p, d)
	m.RemoveDestination(p)
	m.AddDestination(dest)
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
