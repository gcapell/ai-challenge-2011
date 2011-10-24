package main

import (
	"log"
)

type MyBot struct {
	game	*Game
	m	*Map
}

const (
	FOOD_DEPTH    = 7
	EXPLORE_DEPTH = 10
)

//DoTurn is where you should do your bot's actual work.
func (mb *MyBot) DoTurn() {

	// Grab nearby food
	mb.moveToTarget("food", FOOD_DEPTH, func(loc Location) bool { return mb.m.FoodAt(loc) })

	// Attack enemy hill
	mb.moveToTarget("enemy hill", EXPLORE_DEPTH, func(loc Location) bool {return mb.m.EnemyHillAt(loc)})

	// Explore the unknown
	mb.moveToTarget( "explore", EXPLORE_DEPTH, func(loc Location) bool { return !mb.m.squares[loc].wasSeen })

	mb.moveRandomly()
}

// If there's some useful target nearby an ant, move towards it
// Breadth first search
func (mb *MyBot) moveToTarget(reason string, depth int, isTarget func(Location) bool) {
	seen := make(map[Location]bool)

	frontier := make(map[Location]Move) // current -> src

	for loc := range mb.m.MyStationaryAnts() {
		frontier[loc] = Move{src: loc, d: NoMovement}
		seen[loc] = true
	}

	moved := make(map[Location]bool)

	iter := 0
	maxFrontier := 0
	for ; iter < depth; iter++ {

		newFrontier := make(map[Location]Move)
		newMoved := make(map[Location]bool)

		for current, first := range frontier {

			initialMove := first.d == NoMovement

			// Ignore ants who moved last go
			if moved[first.src] || newMoved[first.src] {
				continue
			}

			for d, next := range mb.m.NextValidMoves(current) {
				if seen[next] {
					continue
				}
				if initialMove {
					first.d = d
				}
				seen[next] = true
				if isTarget(next) {
					log.Printf("moving %v for %s\n", first, reason)
					mb.m.IssueOrderLoc(first.src, first.d)
					newMoved[first.src] = true
					break
				} else {
					newFrontier[next] = first
				}
			}
		}
		frontierSize := len(newFrontier)
		if frontierSize == 0 {
			break
		}
		if frontierSize > maxFrontier {
			maxFrontier = frontierSize
		}
		frontier = newFrontier
		moved = newMoved
	}
	log.Printf("Searched to depth %d, width %d for  %s\n", iter, maxFrontier, reason)
}

// Any ants not yet moving should move randomly
func (mb *MyBot) moveRandomly() {
	for loc := range mb.m.MyStationaryAnts() {
		log.Println("randomly moving", loc)
		for d, _ := range mb.m.NextValidMoves(loc) {
			mb.m.IssueOrderLoc(loc, d)
			break
		}
	}
}
