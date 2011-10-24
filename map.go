package main

import (
	"log"
	"strconv"
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

type Map struct {
	Rows int
	Cols int

	itemGrid []Item

	Ants         map[Location]Item
	Dead         map[Location]Item
	Hills        map[Location]Item
	Water        map[Location]bool
	Food         map[Location]bool
	Destinations map[Location]bool
	MyAnts       map[Location]bool // ant location -> is moving?

	viewradius2 int
}

const (
	UNKNOWN Item = iota - 5
	WATER
	FOOD
	LAND
	DEAD
	MY_ANT //= 0
	PLAYER1
	PLAYER2
	PLAYER3
	PLAYER4
	PLAYER5
	PLAYER6
	PLAYER7
	PLAYER8
	PLAYER9
	PLAYER10
	PLAYER11
	PLAYER12
	PLAYER13
	PLAYER14
	PLAYER15
	PLAYER16
	PLAYER17
	PLAYER18
	PLAYER19
	PLAYER20
	PLAYER21
	PLAYER22
	PLAYER23
	PLAYER24
	PLAYER25
)

var (
	DIRS = []Direction{North, East, South, West}
)


//Symbol returns the symbol for the ascii diagram
func (o Item) Symbol() byte {
	switch o {
	case UNKNOWN:
		return '.'
	case WATER:
		return '%'
	case FOOD:
		return '*'
	case LAND:
		return ' '
	case DEAD:
		return '!'
	}

	if o < MY_ANT || o > PLAYER25 {
		log.Panicf("invalid item: %v", o)
	}

	return byte(o) + 'a'
}

//FromSymbol reverses Symbol
func FromSymbol(ch byte) Item {
	switch ch {
	case '.':
		return UNKNOWN
	case '%':
		return WATER
	case '*':
		return FOOD
	case ' ':
		return LAND
	case '!':
		return DEAD
	}
	if ch < 'a' || ch > 'z' {
		log.Panicf("invalid item symbol: %v", ch)
	}
	return Item(ch) + 'a'
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

//NewMap returns a newly constructed blank map.
func NewMap(Rows, Cols, viewradius2 int) *Map {
	m := &Map{
		Rows:        Rows,
		Cols:        Cols,
		Water:       make(map[Location]bool),
		itemGrid:    make([]Item, Rows*Cols),
		viewradius2: viewradius2,
	}
	m.Reset()
	return m
}

//String returns an ascii diagram of the map.
func (m *Map) String() string {
	str := ""
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			s := m.itemGrid[row*m.Cols+col].Symbol()
			str += string([]byte{s}) + " "
		}
		str += "\n"
	}
	return str
}

//Reset clears the map (except for water) for the next turn
func (m *Map) Reset() {
	for i := range m.itemGrid {
		m.itemGrid[i] = UNKNOWN
	}
	for i, val := range m.Water {
		if val {
			m.itemGrid[i] = WATER
		}
	}
	m.Ants = make(map[Location]Item)
	m.Dead = make(map[Location]Item)
	m.Food = make(map[Location]bool)
	m.Destinations = make(map[Location]bool)
	m.Hills = make(map[Location]Item)
	m.MyAnts = make(map[Location]bool)
}

//Item returns the item at a given location
func (m *Map) Item(loc Location) Item {
	return m.itemGrid[loc]
}

func (m *Map) AddWater(loc Location) {
	m.Water[loc] = true
	m.itemGrid[loc] = WATER
}

func (m *Map) AddHill(loc Location, ant Item) {
	m.Hills[loc] = ant
	// m.itemGrid[loc] = ???
}

func (m *Map) AddAnt(loc Location, ant Item) {
	m.Ants[loc] = ant
	m.itemGrid[loc] = ant

	//if it turns out that you don't actually use the visible radius for anything,
	//feel free to comment this out. It's needed for the image debugging, though.
	if ant == MY_ANT {
		m.AddDestination(loc)
		m.AddLand(loc)
		m.MyAnts[loc] = false
	}
}

//AddLand adds a circle of land centered on the given location
func (m *Map) AddLand(center Location) {
	m.DoInRad(center, m.viewradius2, func(row, col int) {
		loc := m.FromRowCol(row, col)
		if m.itemGrid[loc] == UNKNOWN {
			m.itemGrid[loc] = LAND
		}
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

func (m *Map) AddDeadAnt(loc Location, ant Item) {
	m.Dead[loc] = ant
	m.itemGrid[loc] = DEAD
}

func (m *Map) AddFood(loc Location) {
	m.Food[loc] = true
	m.itemGrid[loc] = FOOD
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
	return !m.Water[loc] && !m.Destinations[loc]
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

const (
	North Direction = iota
	East
	South
	West

	NoMovement
)

var directionstrings = map[Direction]string{
	North:      "n",
	South:      "s",
	West:       "w",
	East:       "e",
	NoMovement: "-",
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

func (m *Map) wordsToLoc(words []string) Location {
	if len(words) < 3 {
		log.Panicf("Invalid command format: \"%v\"", words)
	}
	row, _ := strconv.Atoi(words[1])
	col, _ := strconv.Atoi(words[2])
	return m.FromRowCol(row, col)
}

func (m *Map) wordsToAnt(words []string) (Location, Item) {
	if len(words) < 4 {
		log.Panicf("Invalid command format (not enough parameters for ant): \"%v\"", words)
	}
	row, _ := strconv.Atoi(words[1])
	col, _ := strconv.Atoi(words[2])
	loc := m.FromRowCol(row, col)
	ant, _ := strconv.Atoi(words[3])

	return loc, Item(ant)
}

func (m *Map) Update(words []string) {
	switch words[0] {
	case "w":
		m.AddWater(m.wordsToLoc(words))
	case "f":
		m.AddFood(m.wordsToLoc(words))
	case "h":
		m.AddHill(m.wordsToAnt(words))
	case "a":
		m.AddAnt(m.wordsToAnt(words))
	case "d":
		m.AddDeadAnt(m.wordsToAnt(words))
	default:
		log.Panicf("unknown command: %v\n", words)
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

