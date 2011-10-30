package main

import (
	"log"
	"rand"
)

const (
	DEAD_ENEMY_WEIGHT    = 11
	DEAD_FRIENDLY_WEIGHT = -10
)

//DoTurn is where you should do your bot's actual work.
func (m *Map) DoTurn(t *Timer) {
	m.closeCombat()
	t.Split("closeCombat")

	m.defend()
	t.Split("defend")

	m.forage()
	t.Split("forage")

	m.attackEnemyHill()
	t.Split("enemyHill")

	m.scout()
	t.Split("scout")

	m.moveAll()
	t.Split("moveAll")

	t.Split("doneTurn")
}

// Grab any food we know about
func (m *Map) forage() {
	foragers := m.FreeAnts(true)
	for _, assignment := range assign1(foragers, m.food) {
		assignment.ant.moveTo(m, assignment.p, "food")
	}
}

// If there are enemies near our hill, intercept them
func (m *Map) defend() {
	defenders := m.FreeAnts(true)
	for _, assignment := range assign1(defenders, m.EnemiesNearOurHill(VIEWRADIUS2*2)) {
		a, enemy := assignment.ant, assignment.p
		hill := m.nearestHillToDefend(enemy)
		dst := intercept(a.p, enemy, hill)
		assignment.ant.moveTo(m, dst, "intercept")
	}
}

// explore, farm, ...
func (m *Map) scout() {
	scouts := m.FreeAnts(false)
	size := min(ROWS, COLS)
	step := 5

	for _, a := range scouts {
		a.Scout(m, step, size/2)
	}
}

// Attack enemy hill
func (m *Map) attackEnemyHill() {
	if m.targetHill == nil {
		return
	}
	for _, soldier := range m.FreeAnts(true) {
		soldier.moveTo(m, *m.targetHill, "enemy hill")
	}
}

func (a *Ant) Scout(m *Map, step, maxRadius int) {

	targets := spiral(a.p, step, maxRadius)
	for _, p := range targets {
		if m.ShouldExplore(p) {
			a.Explore(m, p)
			log.Printf("%s scouting %v", a, p)
			return
		}
	}
	p := targets[rand.Intn(len(targets))]
	log.Printf("%s randomly scouting %v", a, p)
	a.Explore(m, p)
}

func (a *Ant) Explore(m *Map, p Point) {
	a.moveTo(m, p, "explore")
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

func spiral(p Point, step, maxDistance int) []Point {
	r := make([]Point, 0, maxDistance/step*maxDistance/step)
	for radius := step; radius < maxDistance; radius += step {
		for off := 0; off < radius; off += step {
			r1 := Point{p.r + radius - off, p.c - radius}.sanitised()
			r2 := Point{p.r - radius, p.c - radius + off}.sanitised()
			r3 := Point{p.r - radius + off, p.c + radius}.sanitised()
			r4 := Point{p.r + radius, p.c + radius - step}.sanitised()

			r = append(r, r1, r2, r3, r4)
		}
	}
	return r
}
