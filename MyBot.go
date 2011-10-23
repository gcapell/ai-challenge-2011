package main

import (
	"os"
	"log"
)

type MyBot struct {

}

const (
	FOOD_DEPTH = 7
)

var (
	DIRS = []Direction{North, East, South, West}
)

//NewBot creates a new instance of your bot
func NewBot(s *State) Bot {
	mb := &MyBot{
	//do any necessary initialization here
	}
	return mb
}

//DoTurn is where you should do your bot's actual work.
func (mb *MyBot) DoTurn(s *State) os.Error {

	mb.moveToNearbyFood(s, s.Map)
	mb.moveRandomly(s)
	//returning an error will halt the whole program!
	return nil
}

// If there's food nearby an ant, move towards it
// Breadth first search
func (mb *MyBot) moveToNearbyFood(s *State, m *Map) {
	seen := make(map[Location]bool)

	frontier := make(map[Location]Move)	// current -> src
	
	for loc := range m.MyStationaryAnts() {
		frontier[loc] = Move{src:loc, d:NoMovement}
		seen[loc] = true
	}

	moved := make(map[Location]bool)

	for iter:=0; iter<FOOD_DEPTH; iter++ {
		log.Printf("frontier %d, %v\n", len(frontier), frontier)

		newFrontier := make(map[Location]Move)
		newMoved := make(map[Location]bool)

		for current, first := range frontier {

			initialMove := first.d == NoMovement

			// Ignore ants who moved last go
			if moved[first.src] {
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
				if m.Food[next] {
					log.Printf("moving %v for food\n", first)
					s.IssueOrderLoc(first.src, first.d)
					newMoved[first.src] = true
					break
				} else {
					newFrontier[next] = first
				}
			}
		}
		frontier = newFrontier
		moved = newMoved
	}
}


// Any ants not yet moving should move randomly
func (mb *MyBot) moveRandomly(s *State) {
	for loc := range s.Map.MyStationaryAnts() {
		log.Println("randomly moving", loc)
		for d, _ := range s.Map.NextValidMoves(loc) {
			s.IssueOrderLoc(loc, d)
			break
		}
	}
}