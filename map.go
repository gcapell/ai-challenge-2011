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

	Point struct {x,y int}

	Square struct {
		isWater bool
		wasSeen bool	// Have we ever seen this square?
		lastSeen Turn	// .. if so, when?
	}

	Map struct {
		squares	[]Square

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

func (m *Map) Init(rows, cols, viewRadius2 int) {
	ROWS = rows
	COLS = cols
	VIEWRADIUS2 = viewRadius2

	m.Food = make(map[Location]Turn)
	m.Hills = make(map[Location]Item)

	nSquares := rows * cols
	m.squares = make([]Square, nSquares)
	m.Reset()
}

//Reset clears the map for the next turn
func (m *Map) Reset() {
	m.Ants = make(map[Location]Item)
	m.Destinations = make(map[Location]bool)
	m.MyAnts = make(map[Location]bool)
}

// Given start location, return map of direction -> next location
func (m *Map) NextValidMoves(loc Location) map[Direction]Location {
	next := make(map[Direction]Location)

	for _, d := range DIRS {

		loc2 := m.Move(loc, d)
		if m.SafeDestination(loc2) {
			next[d] = loc2
		}
	}
	return next
}

func (m *Map) EnemyHillAt(loc Location) bool {
	item, found := m.Hills[loc]
	return found && item != MY_ANT
}

func (m *Map) FoodAt(loc Location) bool {
	_, found := m.Food[loc]
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
func (m *Map) ViewFrom(center Location) {
	for _, loc := range(m.Neighbours(center, VIEWRADIUS2)) {
		s := &m.squares[loc]
		s.wasSeen = true
		s.lastSeen = TURN
	}
}

func (m *Map) Neighbours(center Location, rad2 int) [] Location{
	x, y := toRC(center)
	
	reply := make([]Location, 0, rad2)

	for dx:=0; dx*dx<= rad2; dx++ {
		for dy := 0; dy*dy + dx  *dx <= rad2; dy++ {
			reply = append(reply, toLoc(x+dx, y+dy))
			if(dx != 0) {
				reply = append(reply, toLoc(x-dx, y+dy))
			}
			if (dy !=0) {
				reply = append(reply, toLoc(x+dx, y-dy))
			}
			if (dx !=0 && dy != 0) {
				reply = append(reply, toLoc(x-dx, y-dy))
			}
		}
	}
	return reply
}

func (m *Map) AddDestination(loc Location) {
	if m.Destinations[loc] {
		log.Panicf("Already have something at that destination!")
	}
	m.Destinations[loc] = true
}

func (m *Map) RemoveDestination(loc Location) {
	m.Destinations[loc] = false, false
}

//SafeDestination will tell you if the given location is a 
//safe place to dispatch an ant. It considers water and both
//ants that have already sent an order and those that have not.
func (m *Map) SafeDestination(loc Location) bool {
	return !m.squares[loc].isWater && !m.Destinations[loc]
}

func toLoc(row, col int) Location {
	if row < 0 {
		row += ROWS
	}
	if row >= ROWS {
		row -= ROWS
	}
	if col < 0 {
		col += COLS
	}
	if col >= COLS {
		col -= COLS
	}
	return Location(row * COLS + col)
}

func toRC(loc Location) (row, col int) {
	iLoc := int(loc)
	return iLoc / COLS, iLoc % COLS
}

func (d Direction) String() string {
	return directionstrings[d]
}

//Move returns a new location which is one step in the specified direction from the specified location.
func (m *Map) Move(loc Location, d Direction) Location {
	row, col := toRC(loc)
	switch d {
	case North:
		row -= 1
	case South:
		row += 1
	case West:
		col -= 1
	case East:
		col += 1
	case NoMovement: //do nothing
	default:
		log.Panicf("%v is not a valid direction", d)
	}
	return toLoc(row, col)
}

func (m *Map) MarkWater(loc Location) {
	m.squares[loc].isWater = true
}

func (m *Map) MarkFood(loc Location) {
	m.Food[loc] = TURN
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

	loc := toLoc(atoi(words[1]), atoi(words[2]))
	var ant Item
	if len(words) == 4 {
		ant = Item(atoi(words[3]))
	}

	switch words[0] {
	case "w":
		m.MarkWater(loc)
	case "f":
		m.MarkFood(loc)
	case "h":
		m.Hills[loc] = ant
	case "a":
		m.AddAnt(loc, ant)
	case "d":
		// Ignore dead ant
	default:
		log.Panicf("unknown command: %v\n", words)
	}
}

func (m *Map) AddAnt(loc Location, ant Item) {
	m.Ants[loc] = ant

	//if it turns out that you don't actually use the visible radius for anything,
	//feel free to comment this out. It's needed for the image debugging, though.
	if ant == MY_ANT {
		m.AddDestination(loc)
		m.ViewFrom(loc)
		m.MyAnts[loc] = false
	}
}


//Call IssueOrderLoc to issue an order for an ant at loc
func (m *Map) IssueOrderLoc(loc Location, d Direction) {
	dest := m.Move(loc, d)
	m.RemoveDestination(loc)
	m.AddDestination(dest)
	m.MyAnts[loc] = true
	row, col := toRC(loc)
	fmt.Println("o", row, col, d)
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
			loc := toLoc(row, col)
			switch letter {
			case '#':
				// Unknown territory
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

