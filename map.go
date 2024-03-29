package main

import (
	"log"
	"fmt"
	"time"
)

type (
	//Item represents all the various items that may be on the map
	Item int8
	Turn uint

	Square struct {
		isWater  bool
		hasFood  bool
		wasSeen  bool // Have we ever seen this square?
		lastSeen Turn // .. if so, when?
	}

	Ant struct {
		Point  // Where are we now?
		plan     []Point // Where will we be?
		seen     Turn
		reason   string // why are we moving?
		hasMoved	bool	// Has this ant moved this turn?
	}

	Map struct {
		squares [][]Square

		myAnts     map[Location]*Ant
		myHills    []Point
		enemyHills []Point

		enemies []Point
		enemyCombatants []Point	// enemies close to one of our ants
		food    []Point

		hasTargetHill bool
		targetHill    Point // Remember hill we're attacking

		// Places that we're sending ants to already
		exploreTargets map[Location]bool

		thinkTime int64 // thinking time, in nanoseconds
		deadline  int64 // deadline, ns since epoch
		
		// For testing, add a border that we don't draw
		hasBorder bool

		movesThisTurn map[string] []Point
	}
)

const (
	UNKNOWN Item = iota - 5
	WATER
	FOOD
	LAND
	DEAD
	MY_ANT    = 0
	ENEMY_ANT = 1

	MAXPLAYER = 24
	
	BORDER = 4
)

var (
	ROWS, COLS, VIEWRADIUS2, ATTACKRADIUS2 int
	TURN                    Turn
)

func (m *Map) setDeadline() {
	m.deadline = time.Nanoseconds() + m.thinkTime
}

func SetAttackRadius2(ar2 int) {
	ATTACKRADIUS2 = ar2
}

func (m *Map) SetNullDeadline() {
	m.deadline = 0
}
func (m *Map) deadlineExpired() bool {
	if m.deadline > 0 &&  time.Nanoseconds() > m.deadline {
		log.Printf("deadline expired")
		return true
	}
	return false
}

func (m *Map) nearestHillToDefend(p Point) Point {
	var closest Point
	for i, hill := range m.myHills {
		if i == 0 {
			closest = hill
		} else {
			if p.CrowDistance2(hill) < p.CrowDistance2(closest) {
				closest = hill
			}
		}
	}
	return closest
}

func (m *Map) EnemiesNearOurHill(tooClose int) []Point {
	reply := make([]Point, 0)

enemyLoop:
	for _, e := range m.enemies {
		for _, hill := range m.myHills {
			if e.CrowDistance2(hill) <= tooClose {
				reply = append(reply, e)
				continue enemyLoop
			}
		}
	}
	return reply
}

func (a *Ant) String() string {
	if len(a.plan) > 0 {
		return fmt.Sprintf("Ant@%v->%v(%s)", a.Point, a.plan[len(a.plan)-1], a.reason)
	}
	return fmt.Sprintf("Ant@%v", a.Point)
}

func (m *Map) Init(rows, cols, viewRadius2 int, hasBorder bool) {
	if hasBorder {
		rows += BORDER
		cols += BORDER
	}
	ROWS = rows
	COLS = cols
	VIEWRADIUS2 = viewRadius2

	m.myAnts = make(map[Location]*Ant)
	m.myHills = make([]Point, 0)
	m.enemies = make([]Point, 0)
	m.enemyHills = make([]Point, 0)
	m.food = make([]Point, 0)
	m.exploreTargets = make(map[Location]bool)
	m.hasBorder = hasBorder

	m.squares = make([][]Square, rows)
	for row := 0; row < rows; row++ {
		m.squares[row] = make([]Square, cols)
	}
	m.Reset()
}

//Reset clears the map for the next turn
func (m *Map) Reset() {
	m.myHills = m.myHills[:0]
	m.enemies = m.enemies[:0]
	m.enemyHills = m.enemyHills[:0]
	m.enemyCombatants = m.enemyCombatants[:0]
	m.food = m.food[:0]
	m.movesThisTurn = make(map[string] []Point)

	// reset squares
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			s := &m.squares[r][c]
			s.hasFood = false
		}
	}

	// reset ants
	for _, a := range m.myAnts {
		a.hasMoved = false
	}
}

func (m *Map) isWet(p Point) bool {
	return m.squares[p.r][p.c].isWater
}

func (m *Map)  hasFood(p Point) bool {
	return m.squares[p.r][p.c].hasFood
}

// Are we blocked from moving to 'p' ?
func (m *Map) isBlocked(p Point) bool {
	s := &m.squares[p.r][p.c]
	return s.isWater || s.hasFood
}

// Find point one square from 'src' which gets us closer to 'dst' and ok to move to this turn.
func (m *Map) CloserSquare(src, dst Point) (next Point, found bool) {
	path, err := src.ShortestPath(dst, m)
	if err != nil {
		log.Printf("Could not find CloserSquare(%s,%s): %s", src,dst, err)
		return next, false
	}
	next = path[0]
	
	// If there's food there, just wait
	if m.hasFood(next) {
		return next, false
	}
	return path[0], true
}

func (m *Map) AccessibleNeighbours(p Point) []Point {
	allNeighbours := []Point{
		Point{p.r + 1, p.c},
		Point{p.r - 1, p.c},
		Point{p.r, p.c + 1},
		Point{p.r, p.c - 1},
	}
	reply := make([]Point, 4)
	pos := 0
	for _, n := range allNeighbours {
		n.sanitise()
		if m.isWet(n) {
			continue
		}
		if _, ok := m.myAnts[n.loc()]; ok {
			continue
		}
		reply[pos] = n
		pos += 1
	}
	return reply[:pos]
}

// A scout has reported from 'p'
func (m *Map) ViewFrom(scout Point) {
	for _, p := range scout.Neighbours(VIEWRADIUS2) {
		s := &m.squares[p.r][p.c]
		s.wasSeen = true
		s.lastSeen = TURN
	}
}

func (m *Map) MarkWater(p Point) {
	m.squares[p.r][p.c].isWater = true
}

func (m *Map) MarkFood(p Point) {
	m.food = append(m.food, p)
	m.squares[p.r][p.c].hasFood = true
}

func (m *Map) MarkHill(p Point, ant Item) {
	if ant == MY_ANT {
		m.myHills = append(m.myHills, p)
	} else {
		m.enemyHills = append(m.enemyHills, p)
	}
}

func (m *Map) Update(words []string) {
	if words[0] == "turn" {
		turn := Turn(atoi(words[1]))
		if turn != TURN+1 {
			log.Panicf("Turn number out of sync, expected %v got %v", TURN+1, turn)
		}
		TURN = turn
		log.SetPrefix(fmt.Sprintf("%s %d ", *logPrefix, turn))
		return
	}

	p := Point{atoi(words[1]), atoi(words[2])}
	var ant Item
	if len(words) == 4 {
		ant = Item(atoi(words[3]))
	}

	switch words[0] {
	case "w":
		m.MarkWater(p)
	case "f":
		m.MarkFood(p)
	case "h":
		m.MarkHill(p, ant)
	case "a":
		m.AddAnt(p, ant)
	case "d":
		m.DeadAnt(p, ant)
	default:
		log.Panicf("unknown command updating map: %v\n", words)
	}
}

func (m *Map) AddAnt(p Point, ant Item) {
	if ant == MY_ANT {
		m.ViewFrom(p)

		antp, found := m.myAnts[p.loc()]
		if found { // existing ant?
			antp.seen = TURN
		} else { // new ant?
			m.myAnts[p.loc()] = &Ant{Point: p, seen: TURN}
		}
	} else {
		m.enemies = append(m.enemies, p)
	}
}

func (m *Map) DeadAnt(p Point, ant Item) {
	if ant != MY_ANT {
		return
	}
	// FIXME - mark that whatever this ant was trying to do
	// still needs to be done.
	m.myAnts[p.loc()] = nil, false
}

// We have received all the updates for this turn,
// make sure they make sense.
func (m *Map) UpdatesProcessed() {
	// Any ants that missed an update?
	for loc, ant := range m.myAnts {
		if ant.seen != TURN {
			for p, a2 := range m.myAnts {
				log.Printf("ALL ANTS: %v -> %s", p.point(), a2)
			}
			log.Printf("ROWS: %d", ROWS)
			log.Panicf("%v (@ %v) missed an update\n", ant, loc.point())
		}
	}

	// targetHill destroyed?
	if m.hasTargetHill {
		if m.squares[m.targetHill.r][m.targetHill.c].lastSeen == TURN {
			found := false
			for _, p := range m.enemyHills {
				if p.Equals(m.targetHill) {
					found = true
					break
				}
			}
			if !found {
				log.Printf("enemy hill at %v destroyed", m.targetHill)
				m.hasTargetHill = false
				for _, a := range m.myAnts {
					if len(a.plan) > 0 && m.targetHill.Equals(a.plan[len(a.plan)-1]) {
						a.plan = a.plan[:0]
					}
				}
			}
		} else {
			log.Printf("Assuming enemy hill still at %v", m.targetHill)
		}
	}

	// Acquire target
	if !m.hasTargetHill && len(m.enemyHills) != 0 {
		m.hasTargetHill = true
		m.targetHill = m.enemyHills[0]
		log.Printf("Acquired enemy hill at %v", m.targetHill)
	}
}

// Return slice of ants who aren't already assigned
// If interrupt, include ants who already have a plan
func (m *Map) FreeAnts(interrupt bool) []*Ant {
	reply := make([]*Ant, 0, len(m.myAnts))
	for _, ant := range m.myAnts {
		if ant.hasMoved {
			// Already assigned this turn
			continue
		}
		if len(ant.plan) > 0 && !interrupt {
			// Ant already has a plan
			continue
		}
		reply = append(reply, ant)
	}
	return reply
}

func (a *Ant) CancelPlans() {
	a.plan = a.plan[:0]
}
