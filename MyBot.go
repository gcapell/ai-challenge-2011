package main

//DoTurn is where you should do your bot's actual work.
func (m *Map) DoTurn() {
	m.defend()
	m.forage()
	m.moveAll()
}

// Grab any food we know about
func (m *Map) forage() {
	for _, assignment := range assign1(m.FreeAnts(), m.food) {
		assignment.ant.moveTo(m, assignment.p, "food")
	}
}

// If there are enemies near our hill, intercept them
func (m *Map) defend() {
	for _, assignment := range assign1(m.FreeAnts(), m.EnemiesNearOurHill(20)) {
		a, enemy := assignment.ant, assignment.p
		hill := m.nearestHill(enemy)
		dst := intercept(a.p, enemy, hill)
		assignment.ant.moveTo(m, dst, "intercept")
	}
}
