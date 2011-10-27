package main

type MyBot struct {
	game	*Game
	m	*Map
}

//DoTurn is where you should do your bot's actual work.
func (mb *MyBot) DoTurn() {
	mb.m.forage()
	mb.m.moveAll()
}

func (a *Ant) Distance(p Point) int {
	return a.p.Distance(p)
}

func (m *Map) forage() {
	for _, assignment := range assign1(m.FreeAnts(), m.food) {
		assignment.ant.moveTo(m, assignment.p)
	}
}
