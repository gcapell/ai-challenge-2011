package main

import (
	"testing"
)

func TestGroupCombat(t *testing.T) {
	ATTACKRADIUS2 = 5
	DEAD_ENEMY_WEIGHT    = 100
	DEAD_FRIENDLY_WEIGHT = -110

	verifyGroupCombat(t,
		"support a friend",
		`
		.........
		...a.....
		....a..b.
		.........
		.........
		.........
		`, 
		`
		.........
		....a....
		....a..b.
		.........
		.........
		.........
		`,
	)

	verifyGroupCombat(t,
		"run away when outnumbered",
		`....a..b
		 .......b`, 
		`...a...b
		 .......b`, 
 	)

	verifyGroupCombat(t,
		"Attack when we outnumber",
		`....b..a
		 .......a`, 
		`....b.a.
		 ......a.`, 
 	)

	verifyGroupCombat(t,
		"Reject a swap",
		`....b..a
		 ........`, 
		`a...b...
		 ........`, 
 	)
}
func (m *Map) MovesFromMap() (gm, em GroupMove) {
	for _, a := range m.myAnts {
		gm.dst = append(gm.dst, a.p)
	}
	for _, e := range m.enemies {
		em.dst = append(em.dst, e)
	}
	return gm, em
}

func ScoreFromMap(t *testing.T, s string, expected float64) {
	
	m := new(Map)
	m.InitFromString(0,  s)

	gm, em := m.MovesFromMap()
	gm.score(em)
	if gm.evaluated != 1 {
		t.Error("gm.evaluated", gm.evaluated, "!=1")
	}

	if gm.average != expected {
		t.Errorf("got %f, expected %f from map:\n%s\n", gm.average, expected, m)
	}
}

func TestScore(t *testing.T) {
	ATTACKRADIUS2 = 5
	DEAD_ENEMY_WEIGHT    = 11
	DEAD_FRIENDLY_WEIGHT = -10

	tests := []struct{
		s string
		score float64
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
		ScoreFromMap(t, s.s, s.score)
	}
}

func verifyGroupCombat(t *testing.T, reason, initial, final string) {
	m := new(Map)
	m.InitFromString(0, initial)

	combatZones := m.FindCombatZones()
	assert (len(combatZones) == 1, "%v", combatZones)
	cz := combatZones[0]
	bestMove := cz.GroupCombat(m)

	cz.MakeMove(m, bestMove)
	m.moveAll()

	checkMap(t, m, reason, final)
}

