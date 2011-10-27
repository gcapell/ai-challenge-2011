package main
import (
	"sort"
)

type MyBot struct {
	game	*Game
	m	*Map
}

type Assignment struct {
	ant	*Ant
	p	Point
}

const (
	FOOD_DEPTH    = 7
	EXPLORE_DEPTH = 10
)

//DoTurn is where you should do your bot's actual work.
func (mb *MyBot) DoTurn() {
	m := mb.m
	m.forage()
}

func (a *Ant) moveTo(p Point) {
	// FIXME
}

func (a *Ant) Distance(p Point) int {
	return a.p.Distance(p)
}

func (m *Map) forage() {

	for _, assignment := range m.assign1(m.food) {
		assignment.ant.moveTo(assignment.p)
	}
}

type AssignmentSlice []Assignment

func (a Assignment) distance()  int {
	return a.ant.Distance(a.p)
}

func (as AssignmentSlice) Len() int{
	return len(as)
}

func (as AssignmentSlice) Less(i,j int) bool {
	return as[i].distance() < as[j].distance()
}

func (as AssignmentSlice) Swap(i,j int) {
	as[j], as[i] = as[i], as[j]
}

func (as *AssignmentSlice) add(a Assignment) {
	*as = append(*as, a)
}

// Attempt to assign to each target the nearest available ant
func (m *Map) assign1 (targets []Point) [] Assignment {
	
	as := AssignmentSlice(make([]Assignment, 0, len(targets) * len(m.myAnts)))
	
	// Generate all pairings, then sort
	for _, a := range m.myAnts {
		for _, p := range targets {
			as.add( Assignment{a,p})
		}
	}
	sort.Sort(as)
	return as
}
