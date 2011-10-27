package main
import (
	"sort"
	"fmt"
)

type MyBot struct {
	game	*Game
	m	*Map
}

type Assignment struct {
	ant	*Ant
	p	Point
}

func (a Assignment) String()string {
	return fmt.Sprintf("Assignment{ %v->%v}", a.ant.p, a.p)
}

//DoTurn is where you should do your bot's actual work.
func (mb *MyBot) DoTurn() {
	mb.m.forage()
}

func (a *Ant) moveTo(p Point) {
	// FIXME
}

func (a *Ant) Distance(p Point) int {
	return a.p.Distance(p)
}

func (m *Map) forage() {

	for _, assignment := range assign1(m.FreeAnts(), m.food) {
		assignment.ant.moveTo(assignment.p)
	}
}

type AssignmentSlice []Assignment

// Implement sort.Interface
func (a Assignment) distance()  int {return a.ant.Distance(a.p)}
func (as AssignmentSlice) Len() int{return len(as)}
func (as AssignmentSlice) Less(i,j int) bool {return as[i].distance() < as[j].distance()}
func (as AssignmentSlice) Swap(i,j int) {as[j], as[i] = as[i], as[j]}

func (as *AssignmentSlice) add(a Assignment) {
	*as = append(*as, a)
}

// Return optimal slice of ant->target assignments
// See wikipedia.org/wiki/Assignment_problem
func assign1 (ants []*Ant, targets []Point) [] Assignment {
	
	// For now, we'll be greedy-stupid.
	// Generate all pairings, then sort
	as := AssignmentSlice(make([]Assignment, 0, len(ants) * len(targets)))
	for _, a := range ants {
		if a.isBusy {
			continue
		}
		for _, p := range targets {
			as.add( Assignment{a,p})
		}
	}
	sort.Sort(as)

	var reply [] Assignment

	assigned := make(map[Location]bool)

	for _, a := range as {
		if a.ant.isBusy {
			continue
		}
		if assigned[a.p.loc()] {
			continue
		}
		a.ant.isBusy = true
		assigned[a.p.loc()] = true
		reply = append(reply, a)
	}
	return reply
}
