package main

import (
	"fmt"
)

type Direction int

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

func (d Direction) String() string {
	return directionstrings[d]
}

//Call IssueOrderLoc to issue an order for an ant at loc
func (m *Map) IssueOrderLoc(p Point, d Direction) {
	//...
	fmt.Println("o", p.r, p.c, d)
}

