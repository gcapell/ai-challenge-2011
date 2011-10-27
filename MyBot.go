package main

type MyBot struct {
	game	*Game
	m	*Map
}

//DoTurn is where you should do your bot's actual work.
func (mb *MyBot) DoTurn() {
	mb.m.defend()
	mb.m.forage()
	mb.m.moveAll()
}

func (a *Ant) Distance(p Point) int {
	return a.p.Distance(p)
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

		var dst Point

		if a.p.intercepts(hill, enemy) {
			dst = enemy
		} else {
			dst = between(hill, enemy)
		}

		assignment.ant.moveTo(m, dst, "intercept")
	}
}