package main

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
	mb.moveToTarget("food", FOOD_DEPTH, func(p Point) bool { return mb.m.FoodAt(p) })

	// Attack enemy hill
	mb.moveToTarget("enemy hill", EXPLORE_DEPTH, func(p Point) bool {return mb.m.EnemyHillAt(p)})

	// Explore the unknown
	mb.moveToTarget( "explore", EXPLORE_DEPTH, func(p Point) bool { return !mb.m.squares[p.r][p.c].wasSeen })

	mb.moveRandomly()
}

// If there's some useful target nearby an ant, move towards it
// Breadth first search
func (mb *MyBot) moveToTarget(reason string, depth int, isTarget func(Point) bool) {}

// Any ants not yet moving should move randomly
func (mb *MyBot) moveRandomly() {}
