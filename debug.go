package main

import (
	"log"
)

//Symbol returns the symbol for the ascii diagram
func (o Item) Symbol() byte {
	switch o {
	case UNKNOWN, LAND:
		return '.'
	case WATER:
		return '%'
	case FOOD:
		return '*'
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
	if _, ok := m.myAnts[p.loc()]; ok {
		return MY_ANT
	}
	for _, e := range m.enemies {
		if e.Equals(p) {
			return ENEMY_ANT
		}
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
