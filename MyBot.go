package main

import (
	"log"
	"rand"
	"strings"
	"time"
	"fmt"
)

const (
	// When we have lots of ants, directing them _all_ to start attacking in the one
	// turn takes too long. Instead, direct HILL_ATTACKERS per turn.
	HILL_ATTACKERS = 20
)

//DoTurn is where you should do your bot's actual work.
func (m *Map) DoTurn() {

	strategies := []struct {
		fn   func()
		name string
	}{
		{func() { m.closeCombat() }, "closeCombat"},
		{func() { m.defend() }, "defend"},
		{func() { m.reinforce() }, "reinforce"},
		{func() { m.forage() }, "forage"},
		{func() { m.attackEnemyHill() }, "enemyHill"},
		{func() { m.scout() }, "scout"},
	}
	times := make([]string, 0, len(strategies))

	for _, s := range strategies {
		if m.deadlineExpired() {
			break
		}
		start := time.Nanoseconds()
		s.fn()
		delta_ms := float64(time.Nanoseconds() - start) / 1e6 
		if delta_ms > 100 {
			times = append(times, fmt.Sprintf("%s %.2f", s.name, delta_ms))
		}
	}
	m.moveAll()
	
	log.Print("timings: %s", strings.Join(times, ", "))
}

// Grab any food we know about
func (m *Map) forage() {
	foragers := m.FreeAnts(true)
	for _, assignment := range assign1(foragers, m.food) {
		if m.deadlineExpired() {
			break
		}
		assignment.ant.SetPathFor(m, assignment.p, "food")
	}
}

// If there is a combat zone near us, move towards it
func (m *Map) reinforce() {
	reinforcers := m.FreeAnts(true)
	for _, assignment := range assignNearbyCrow2(reinforcers, m.enemyCombatants, ATTACKRADIUS2 *4) {
		if m.deadlineExpired() {
			break
		}
		next, ok := m.CloserSquare(assignment.ant.Point, assignment.p)
		if ok {
			assignment.ant.OverridingMove(m, next, "reinforce")
		}
	}
}

// If there are enemies near our hill, intercept them
func (m *Map) defend() {
	defenders := m.FreeAnts(true)
	for _, assignment := range assign1(defenders, m.EnemiesNearOurHill(VIEWRADIUS2*2)) {
		if m.deadlineExpired() {
			break
		}
		a, enemy := assignment.ant, assignment.p
		hill := m.nearestHillToDefend(enemy)
		dst := a.intercept(enemy, hill)
		assignment.ant.SetPathFor(m, dst, "intercept")
	}
}

// explore, farm, ...
func (m *Map) scout() {
	scouts := m.FreeAnts(false)
	size := min(ROWS, COLS)
	step := 5
	
	if size < 20 {
		// tiny testing map
		return
	}

	for _, a := range scouts {
		if m.deadlineExpired() {
			break
		}
		a.Scout(m, step, size/2)
	}
}

// Attack enemy hill
func (m *Map) attackEnemyHill() {
	if !m.hasTargetHill {
		return
	}
	log.Printf("atttacking enemy hill at %v, enemyHills:%v, myHills:%v", m.targetHill, m.enemyHills, m.myHills)

	for _, soldier := range m.FreeAnts(true) {
		if m.deadlineExpired() {
			break
		}
		soldier.SetPathFor(m, m.targetHill, "enemy hill")
	}
}

func (a *Ant) Scout(m *Map, step, maxRadius int) {
	targets := a.spiral(step, maxRadius)
	for _, p := range targets {
		if m.ShouldExplore(p) {
			a.Explore(m, p)
			// log.Printf("%s scouting %v", a, p)
			return
		}
	}
	p := targets[rand.Intn(len(targets))]
	a.Explore(m, p)
}

func (a *Ant) Explore(m *Map, p Point) {
	a.SetPathFor(m, p, "explore")
	m.MarkExploreTarget(p)
}

func (m *Map) ShouldExplore(p Point) bool {
	if m.squares[p.r][p.c].wasSeen {
		return false
	}
	for t, _ := range m.exploreTargets {
		if t.point().CrowDistance2(p) < VIEWRADIUS2 {
			return false
		}
	}
	return true
}

func (m *Map) MarkExploreTarget(p Point) {
	m.exploreTargets[p.loc()] = true
}

