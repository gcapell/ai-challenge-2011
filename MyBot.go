package main

import "log"

//DoTurn is where you should do your bot's actual work.
func (m *Map) DoTurn() {
	m.defend()
	m.forage()
	m.scout()
	m.moveAll()
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
	for _, assignment := range assign1(defenders, m.EnemiesNearOurHill(20)) {
		a, enemy := assignment.ant, assignment.p
		hill := m.nearestHill(enemy)
		dst := intercept(a.p, enemy, hill)
		assignment.ant.moveTo(m, dst, "intercept")
	}
}

// explore, farm, ...
func (m *Map) scout() {
	scouts := m.FreeAnts(false)
	log.Println("scouts:", scouts)
}