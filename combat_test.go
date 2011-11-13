package main

import (
	"testing"
	"log"
)


func TestSpeedGroupCombat(t *testing.T) {
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

func checkZone(t *testing.T, z *CombatZone, friendly, enemy int) {
	if len(z.friendly) != friendly {
		t.Fatalf("expected %d friendlies, got %d", friendly, len(z.friendly))
	}
	if len(z.enemy) != enemy {
		t.Fatalf("expected %d ememies, got %d", enemy, len(z.enemy))
	}
}

func checkCombatZones(t *testing.T, s string, friendly, enemy int) {
	m := new(Map)
	m.InitFromString( s, true)
	zones := m.FindCombatZones()
	if len(zones) != 1 {
		t.Fatalf("Expected 1 combat zone got %d in \n%s", len(zones), m)
	}
	checkZone(t, zones[0], friendly, enemy)
}

func TestReinforce(t *testing.T) {
	map1 := `
	.a..b.
	......
	a.....
	....a.
	......
	`
	m := new(Map)
	m.InitFromString( map1, true)
	zones := m.FindCombatZones()
	if len(zones) != 1 {
		log.Fatalf("Expected 1 combat zone got %d in \n%s", len(zones), m)
	}
	checkZone(t, zones[0], 2, 1)

	m.DoTurn()
	checkMap(t, m, "reinforce", `
	a...b.
	......
	.a....
	......
	....a.
	`)
}

func TestFindCombatZones(t *testing.T) {
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

	SimultaneousOverridingMove(cz.friendly, bestMove.dst, m, "combat")
	m.moveAll()

	checkMap(t, m, reason, final)
}
