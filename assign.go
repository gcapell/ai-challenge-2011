package main
import (
	"sort"
	"fmt"
)

type (
	Assignment struct {
		ant	*Ant
		p	Point
	}
	AssignmentSlice []Assignment
) 

func (a Assignment) String()string {
	return fmt.Sprintf("Assignment{ %v->%v}", a.ant.p, a.p)
}

// Implement sort.Interface
func (a Assignment) distance()  int {return a.ant.Distance(a.p)}
func (as AssignmentSlice) Len() int{return len(as)}
func (as AssignmentSlice) Less(i,j int) bool {return as[i].distance() < as[j].distance()}
func (as AssignmentSlice) Swap(i,j int) {as[j], as[i] = as[i], as[j]}

func (as *AssignmentSlice) add(a Assignment) {
	*as = append(*as, a)
}

// Return optimal slice of ant->target assignments
// FIXME: See wikipedia.org/wiki/Assignment_problem
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
