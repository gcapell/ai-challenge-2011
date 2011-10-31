package main

import (
	"testing"
)

func (m *Map) MovesFromMap() (gm, em GroupMove) {
	for _, a := range m.myAnts {
		gm.dst = append(gm.dst, a.p)
	}
	for _, e := range m.enemies {
		em.dst = append(em.dst, e)
	}
	return gm, em
}

func ScoreFromMap(s string, expected int) {
	
	m := new(Map)
	m.InitFromString(0,  s)

	gm, em := m.MovesFromMap()
	gm.score(em)
	assert (gm.evaluated == 1, "")

	assert (gm.best == expected, "got %d, expected %d from map:\n%s\n", gm.best, expected, m)
}

func TestScore(t *testing.T) {
	ATTACKRADIUS2 = 5
	DEAD_ENEMY_WEIGHT    = 11
	DEAD_FRIENDLY_WEIGHT = -10

	tests := []struct{
		s string
		score int
	}{
		{`...b.a
		  ...b.a
		  ...b..
		  ......`, 2},

		{`..a..b
		  .....b`, 0},

		{`...a.b
		  .....b`, -10},

		{`...b.a
		  .....a`, 11},

		{`...b.a
		  ......`, 1},
	}
	for _, s := range tests {
		ScoreFromMap(s.s, s.score)
	}
}
