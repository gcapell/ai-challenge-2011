package main

import (
	"testing"
	"log"
)


func TestSpeedGroupCombat(t *testing.T) {
	ATTACKRADIUS2 = 5

	verifyGroupCombat(t,
		"many units",
		`
		..........
		....a.a...
		..........
		..........
		.a...bb...
		..........
		..........
		b....a....
		..........
		`,
		`
		....a.a...
		..........
		..........
		.a........
		.....bb...
		..........
		..........
		b.........
		.....a....
		`,
	)
}

func checkCombatZones(t *testing.T, s string, friendly, enemy int) {
	m := new(Map)
	m.InitFromString( s, true)
	zones := m.FindCombatZones()
	if len(zones) != 1 {
		log.Fatalf("Expected 1 combat zone got %d in \n%s", len(zones), m)
	}
	z := zones[0]
	if len(z.friendly) != friendly {
		log.Fatalf("expected %d friendlies, got %d", friendly, len(z.friendly))
	}
	if len(z.enemy) != enemy {
		log.Fatalf("expected %d ememies, got %d", enemy, len(z.enemy))
	}
	
}
func TestFindCombatZones(t *testing.T) {
	ATTACKRADIUS2 = 5
	checkCombatZones(t, `
	bbbbb
	.....
	.....
	.aaa.
	`, 3,5 )

	checkCombatZones(t, `
	aaaaa
	.....
	.....
	.bbb.
	`, 5, 3)
}

func TestLargeGroupCombat(t *testing.T) {
	ATTACKRADIUS2 = 5
	
	verifyGroupCombat(t, "charge", 
	`
	aaaaa
	.....
	.....
	.bbb.
	`,`
	.....
	aaaaa
	.....
	.bbb.
	`)
	verifyGroupCombat(t, "flee", 
	`
	bbbbb
	.....
	.....
	.aaa.
	.....
	`,`
	bbbbb
	.....
	.....
	.....
	.aaa.
	`)
	
	verifyGroupCombat(t, "charge when surrounding",
	`
	aaaa
	....
	....
	.bb.
	....
	....
	aaaa
	`,`
	....
	aaaa
	....
	.bb.
	....
	aaaa
	....
	`)
}

func TestGroupCombat(t *testing.T) {
	ATTACKRADIUS2 = 5

	log.Printf("hello\n")
	verifyGroupCombat(t,
		"Attack when we outnumber",
		`
		 ...a
		 b..a
		 ...a`,
		`
		 ..a.
		 b.a.
		 ..a.`,
	)
	log.Printf("goodbye\n")


	verifyGroupCombat(t,
		"support a friend",
		`
		a....
		.a..b
		`,
		`
		.a...
		.a..b
		`,
	)

	verifyGroupCombat(t,
		"run away when outnumbered",
		`.a..b
		 ....b`,
		`a...b
		 ....b`,
	)

	verifyGroupCombat(t,
		"Reject a swap",
		`b..a.
		`,
		`b...a
		 `,
	)
}
func (m *Map) MovesFromMap() (gm, em GroupMove) {
	for _, a := range m.myAnts {
		gm.dst = append(gm.dst, a.Point)
	}
	for _, e := range m.enemies {
		em.dst = append(em.dst, e)
	}
	return gm, em
}

func ScoreFromMap(t *testing.T, s string, expected float64) {

	m := new(Map)
	m.InitFromString(s, true)

	gm, em := m.MovesFromMap()
	gm.score(em, NEAR_OUR_HILL_SCORING)
	if gm.evaluated != 1 {
		t.Error("gm.evaluated", gm.evaluated, "!=1")
	}

	if gm.average != expected {
		t.Errorf("got %f, expected %f from map:\n%s\n", gm.average, expected, m)
	}
}

func TestScore(t *testing.T) {
	ATTACKRADIUS2 = 5

	tests := []struct {
		s     string
		score float64
	}{
		{`...b.a
		  ...b.a
		  ...b..
		  ......`, 20},

		{`..a..b
		  .....b`, 0},

		{`...a.b
		  .....b`, -90},

		{`...b.a
		  .....a`, 100},

		{`...b.a
		  ......`, 10},
	}
	for _, s := range tests {
		ScoreFromMap(t, s.s, s.score)
	}
}

func verifyGroupCombat(t *testing.T, reason, initial, final string) {
	m := new(Map)
	m.InitFromString( initial, true)

	combatZones := m.FindCombatZones()
	assert(len(combatZones) == 1, "%v", combatZones)
	cz := combatZones[0]
	bestMove := cz.GroupCombat(m)

	MakeMove(cz.friendly, bestMove.dst, m)
	m.moveAll()

	checkMap(t, m, reason, final)
}
