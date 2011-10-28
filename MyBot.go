package main

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
	for _, assignment := range assign1(defenders, m.EnemiesNearOurHill(VIEWRADIUS2*2)) {
		a, enemy := assignment.ant, assignment.p
		hill := m.nearestHill(enemy)
		dst := intercept(a.p, enemy, hill)
		assignment.ant.moveTo(m, dst, "intercept")
	}
}

func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}

// explore, farm, ...
func (m *Map) scout() {
	scouts := m.FreeAnts(false)

	for _, a := range scouts {
		size :=min(ROWS, COLS)
		step := 10
		a.Scout(m, step, size/3)
	}
}

func (a *Ant) Scout (m *Map, step, maxRadius int) {
	for _, p := range spiral(a.p, step, maxRadius) {
		if m.ShouldExplore(p) {
			a.moveTo(m, p, "explore")
			m.MarkExploreTarget(p)
		}
	}
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

func (m *Map) MarkExploreTarget (p Point)  {
	m.exploreTargets[p.loc()] = true
}

func spiral(p  Point, step, maxDistance int) []Point {
	r := make ([]Point, maxDistance/step *maxDistance/step)
	for radius := step; radius < maxDistance; radius += step {
		for off := 0; off < radius; off += step {
			r = append(r, 
				Point{p.r + radius -off, p.c - radius}.sanitised(),
				Point{p.r - radius, p.c - radius + off }.sanitised(),
				Point{p.r - radius + off , p.c + radius}.sanitised(),
				Point{p.r + radius , p.c + radius - step}.sanitised(),
			)
		}
	}
	return r
}
