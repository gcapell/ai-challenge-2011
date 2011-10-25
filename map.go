package main

import (
	"log"
	"fmt"
)

//Item represents all the various items that may be on the map
type Item int8

type Move struct {
	src Location
	d   Direction
}

//Location combines (Row, Col) coordinate pairs for use as keys in maps (and in a 1d array)
type Location int

//Direction represents the direction concept for issuing orders.
type Direction int

type Square struct {
	isWater bool
	wasSeen bool	// Have we ever seen this square?
	lastSeen Turn	// .. if so, when?
}

type Map struct {
	Rows int
	Cols int

	squares	[]Square

	Ants         map[Location]Item
	Hills        map[Location]Item
	Food         map[Location]Turn
	Destinations map[Location]bool
	MyAnts       map[Location]bool // ant location -> is moving?

	viewradius2 int

	game	*Game
}

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
	DIRS = []Direction{North, East, South, West}
	directionstrings = map[Direction]string{
		North:      "n",
		South:      "s",
		West:       "w",
		East:       "e",
		NoMovement: "-",
	}
)

func (m *Map) Init(g *Game) {
	m.Rows = g.Rows
	m.Cols = g.Cols
	m.game = g

	m.Food = make(map[Location]Turn)
	m.Hills = make(map[Location]Item)

	nSquares := m.Rows * m.Cols
	m.squares = make([]Square, nSquares)
	m.viewradius2 = g.ViewRadius2
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
	m.DoInRad(center, m.viewradius2, func(row, col int) {
		loc := m.FromRowCol(row, col)
		s := &m.squares[loc]
		s.wasSeen = true
		s.lastSeen = m.game.turn
	})
}

func (m *Map) DoInRad(center Location, rad2 int, Action func(row, col int)) {
	row1, col1 := m.FromLocation(center)
	for row := row1 - m.Rows/2; row < row1+m.Rows/2; row++ {
		for col := col1 - m.Cols/2; col < col1+m.Cols/2; col++ {
			row_delta := row - row1
			col_delta := col - col1
			if row_delta*row_delta+col_delta*col_delta < rad2 {
				Action(row, col)
			}
		}
	}
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

//FromRowCol returns a Location given an (Row, Col) pair
func (m *Map) FromRowCol(Row, Col int) Location {
	for Row < 0 {
		Row += m.Rows
	}
	for Row >= m.Rows {
		Row -= m.Rows
	}
	for Col < 0 {
		Col += m.Cols
	}
	for Col >= m.Cols {
		Col -= m.Cols
	}

	return Location(Row*m.Cols + Col)
}

//FromLocation returns an (Row, Col) pair given a Location
func (m *Map) FromLocation(loc Location) (int, int) {
	iLoc := int(loc)
	return iLoc / m.Cols, iLoc % m.Cols
}

func (d Direction) String() string {
	return directionstrings[d]
}

//Move returns a new location which is one step in the specified direction from the specified location.
func (m *Map) Move(loc Location, d Direction) Location {
	Row, Col := m.FromLocation(loc)
	switch d {
	case North:
		Row -= 1
	case South:
		Row += 1
	case West:
		Col -= 1
	case East:
		Col += 1
	case NoMovement: //do nothing
	default:
		log.Panicf("%v is not a valid direction", d)
	}
	return m.FromRowCol(Row, Col) //this will handle wrapping out-of-bounds numbers
}

func (m *Map) Update(words []string) {
	loc := m.FromRowCol(atoi(words[1]), atoi(words[2]))
	var ant Item
	if len(words) == 4 {
		ant = Item(atoi(words[3]))
	}

	switch words[0] {
	case "w":
		m.squares[loc].isWater = true
	case "f":
		m.Food[loc] = m.game.turn
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
	row, col := m.FromLocation(loc)
	fmt.Println("o", row, col, d)
}

