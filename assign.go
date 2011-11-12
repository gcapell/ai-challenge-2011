package main

import (
	"sort"
	"fmt"
)

type (
	Assignment struct {
		ant      *Ant
		p        Point
		distance int
	}
	AssignmentSlice []Assignment
)

func (a Assignment) String() string {
	return fmt.Sprintf("Assignment{ %v->%v (%v)}", a.ant.Point, a.p, a.distance)
}

// Implement sort.Interface
func (as AssignmentSlice) Len() int           { return len(as) }
func (as AssignmentSlice) Less(i, j int) bool { return as[i].distance < as[j].distance }
func (as AssignmentSlice) Swap(i, j int)      { as[j], as[i] = as[i], as[j] }

func allPairings(ants[]*Ant, targets[]Point, distance func(*Ant, Point) int) AssignmentSlice{
	as := AssignmentSlice(make([]Assignment, len(ants)*len(targets)))
	pos := 0
	for _, a := range ants {
		for _, p := range targets {
			d := distance(a,p)
			if d>0 {
				as[pos] = Assignment{a, p, d}
				pos += 1
			}
		}
	}
	return as[:pos]
}

func filterAssignments(assignments[]Assignment, pointsUsed func(a Assignment) []Point) []Assignment {
	reply := make([]Assignment, len(assignments))

	assigned := make(map[Location]bool)
	pos := 0

	assignedLoop:
	for _, a := range assignments {
		used := pointsUsed(a)
		for _, p := range used {
			if assigned[p.loc()] {
				continue assignedLoop
			}
		}
		for _, p := range used {
			assigned[p.loc()] = true
		}
		reply[pos] = a
		pos += 1
	}
	return reply[:pos]
}

func antAndPoint(a Assignment) []Point {
	return []Point {a.ant.Point, a.p}
}

func pointOnly(a Assignment) []Point {
	return []Point {a.p}
}

// Return optimal slice of ant->target assignments
// FIXME: See wikipedia.org/wiki/Assignment_problem
func assign1(ants []*Ant, targets []Point) []Assignment {

	// For now, we'll be greedy-stupid.
	// Generate all pairings, then sort
	as := allPairings(ants, targets, func(a *Ant, p Point) int {return a.CrowDistance2(p)})
	sort.Sort(as)
	return filterAssignments(as, antAndPoint)
}


// Assign each ant to the nearest target
// (Skip assignments where crow distance squared> maxCrow2)
func assignNearbyCrow2(ants []*Ant, targets [] Point, maxCrow2 int) []Assignment {
	as := allPairings(ants, targets, func(a *Ant, p Point) int {
		d := a.CrowDistance2(p)
		if d > maxCrow2 {
			return 0
		}
		return d
	})
	sort.Sort(as)
	return filterAssignments(as, pointOnly)
}
