package main

import (
	"log"
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

	if o < MY_ANT || o > MAXPLAYER {
		log.Panicf("invalid item: %v", o)
	}

	return byte(o) + 'a'
}

func (m *Map) ItemAt(loc Location) Item {
	s := &m.squares[loc]
	if s.isWater {
		return WATER
	}
	ant, found := m.Ants[loc]
	if found {
		return ant
	}
	_, found = m.Food[loc]
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
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			loc := m.FromRowCol(row, col)
			s := m.ItemAt(loc).Symbol()
			str += string([]byte{s})
		}
		str += "\n"
	}
	return str
}
