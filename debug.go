package main

import (
	"log"
)

//Symbol returns the symbol for the ascii diagram
func (o Item) Symbol() byte {
	switch o {
	case UNKNOWN:
		return '#'
	case WATER:
		return '%'
	case FOOD:
		return '*'
	case LAND:
		return '.'
	case DEAD:
		return '!'
	}

	if o < MY_ANT || o > MAXPLAYER {
		log.Panicf("invalid item: %v", o)
	}

	return byte(o) + 'a'
}

func (m *Map) ItemAt(p Point) Item {
	s := &m.squares[p.r][p.c]
	if s.isWater {
		return WATER
	}
	ant, found := m.Ants[p.loc()]
	if found {
		return ant
	}
	_, found = m.Food[p.loc()]
	if found {
		return FOOD
	}
	if !s.wasSeen {
		return UNKNOWN
	}

	return LAND
}

//String returns an ascii diagram of the map.
func (m *Map) String() string {
	str := ""
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			s := m.ItemAt(Point{row, col}).Symbol()
			str += string([]byte{s})
		}
		str += "\n"
	}
	return str
}
