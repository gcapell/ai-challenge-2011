package main

import (
	"testing"
	"log"
)

func TestGroupCombat(t *testing.T) {
	ATTACKRADIUS2 = 5
	DEAD_ENEMY_WEIGHT    = 11
	DEAD_FRIENDLY_WEIGHT = -10

	log.Printf("Testing Group Combat")
	verifyGroupCombat(t,
		"run away when outnumbered",
		`....a..b
		 .......b`, 
		`...a...b
		 .......b`, 
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

func ScoreFromMap(t *testing.T, s string, expected int) {
	
	m := new(Map)
	m.InitFromString(0,  s)

	gm, em := m.MovesFromMap()
	gm.score(em)
	if gm.evaluated != 1 {
		t.Error("gm.evaluated", gm.evaluated, "!=1")
	}

	if gm.best != expected {
		t.Errorf("got %d, expected %d from map:\n%s\n", gm.best, expected, m)
	}
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
		ScoreFromMap(t, s.s, s.score)
	}
}

func verifyGroupCombat(t *testing.T, reason, initial, final string) {
	m := new(Map)
	m.InitFromString(0, initial)

	combatZones := m.FindCombatZones()
	assert (len(combatZones) == 1, "%v", combatZones)
	cz := combatZones[0]
	log.Printf("cz: %+v", cz)
	bestMove := cz.GroupCombat(m)
	log.Printf("bestMove: %+v", bestMove)

	cz.MakeMove(m, bestMove)
	log.Printf("m.myAnts: %+v", m.myAnts)
	m.moveAll()
	log.Printf("after move: %+v", m.myAnts)

	log.Printf("post combat:\n%s", m)

	checkMap(t, m, reason, final)
}

