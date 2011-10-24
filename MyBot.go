package main

import (
	"os"
	"log"
)

type MyBot struct {

}

const (
	FOOD_DEPTH    = 7
	EXPLORE_DEPTH = 10
)

var (
	DIRS = []Direction{North, East, South, West}
)

//NewBot creates a new instance of your bot
func NewBot(s *Game) Bot {
	mb := &MyBot{
	//do any necessary initialization here
	}
	return mb
}

//DoTurn is where you should do your bot's actual work.
func (mb *MyBot) DoTurn(s *Game) os.Error {

	// Grab nearby food
	mb.moveToTarget(s, s.Map, "food", FOOD_DEPTH, func(loc Location) bool { return s.Map.Food[loc] })

	// Attack enemy hill
	mb.moveToTarget(s, s.Map, "enemy hill", EXPLORE_DEPTH, func(loc Location) bool {
		if item, found := s.Map.Hills[loc]; found {
			if item != MY_ANT {
				return true
			}
		}
		return false
	})

	// Explore the unknown
	mb.moveToTarget(s, s.Map, "explore", EXPLORE_DEPTH, func(loc Location) bool { return s.Map.itemGrid[loc] == UNKNOWN })

	mb.moveRandomly(s)
	//returning an error will halt the whole program!
	return nil
}
// If there's some useful target nearby an ant, move towards it
// Breadth first search
func (mb *MyBot) moveToTarget(s *Game, m *Map, reason string, depth int, isTarget func(Location) bool) {
	seen := make(map[Location]bool)

	frontier := make(map[Location]Move) // current -> src

	for loc := range m.MyStationaryAnts() {
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

			for d, next := range m.NextValidMoves(current) {
				if seen[next] {
					continue
				}
				if initialMove {
					first.d = d
				}
				seen[next] = true
				if isTarget(next) {
					log.Printf("moving %v for %s\n", first, reason)
					s.IssueOrderLoc(first.src, first.d)
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
func (mb *MyBot) moveRandomly(s *Game) {
	for loc := range s.Map.MyStationaryAnts() {
		log.Println("randomly moving", loc)
		for d, _ := range s.Map.NextValidMoves(loc) {
			s.IssueOrderLoc(loc, d)
			break
		}
	}
}
