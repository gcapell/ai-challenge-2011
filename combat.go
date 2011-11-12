package main

import (
	"log"
	"fmt"
	"strings"
)

type (
	CombatZone struct {
		zone     int
		friendly []Point
		enemy    []Point
	}
	GroupMove struct {
		dst          []Point
		updated      bool
		worst, total int
		average      float64
		evaluated    int
	}
	ScoringHeuristic struct {
		deadEnemy, deadFriendly int
	}
)

func (gm *GroupMove) String() string {
	var u string
	if gm.updated {
		u = "U"
	} else {
		u = "u"
	}
	return fmt.Sprintf("GM{ %v %s %d/%.1f after %d}",
		gm.dst, u, gm.worst, gm.average, gm.evaluated)
}

var (
	NEAR_OUR_HILL_SCORING = ScoringHeuristic{deadEnemy: 100, deadFriendly: -90}
	SCOUTING_SCORING      = ScoringHeuristic{deadEnemy: 100, deadFriendly: -110}
)

// Minimax for close combat
func (m *Map) closeCombat() {
	zones := m.FindCombatZones()
	messages := make([]string,0,len(zones))
	for _, cz := range zones {
		if m.deadlineExpired() {
			break
		}
		bestMove := cz.GroupCombat(m)
		messages = append(messages, fmt.Sprintf("%d/%d", len(cz.friendly), len(cz.enemy)))
		if bestMove != nil {
			MakeMove(cz.friendly, bestMove.dst, m)
		}
	}
	log.Printf("group combat %s", strings.Join(messages, ", "))
}

func MakeMove(src, dst []Point, m *Map) {
	for i, srcP := range src {
		dstP := dst[i]
		ant := m.myAnts[srcP.loc()]
		ant.moveToPoint(m, dstP, "combat")
	}
}

func (m *Map) FindCombatZones() []*CombatZone {
	zoneNum = 0
	zones := make([]*CombatZone, 0)

	for _, e := range m.enemies {
		myZone := NewZone(e)
		merging := false
		for _, a := range m.FriendliesInRangeOf(e) {
			// log.Printf("%v in range of %v", a, e)
			found := false
			// Is this friendly part of an existing zone?
			for _, z := range zones {
				if a.In(z.friendly) {
					found = true
					if !merging {
						merging = true
						z.enemy = append(z.enemy, e)
						z.friendly = append(z.friendly, myZone.friendly...)
						myZone = z
					} else {
						z.zone = myZone.zone
					}
				}
			}
			if !found {
				myZone.friendly = append(myZone.friendly, a)
			}
		}
		if !merging && len(myZone.friendly) != 0 {
			zones = append(zones, myZone)
		}
		if false {
			log.Printf("zones")
			for _, z := range zones {
				log.Printf("zone: %+v", z)
			}
		}
	}
	// Final merge
	seen := make(map[int]bool)
	reply := make([]*CombatZone, 0, len(zones))
	for i, cz := range zones {
		if seen[cz.zone] {
			continue
		}
		seen[cz.zone] = true
		for _, otherCz := range zones[i+1:] {
			if otherCz.zone == cz.zone {
				cz.friendly = append(cz.friendly, otherCz.friendly...)
				cz.enemy = append(cz.enemy, otherCz.enemy...)
			}
		}
		reply = append(reply, cz)
	}
	return reply
}

func (cz *CombatZone) GetScoringHeuristic(m *Map) ScoringHeuristic {
	for _, p := range cz.enemy {
		for _, h := range m.myHills {
			if p.CrowDistance2(h) <= 2*VIEWRADIUS2 {
				return NEAR_OUR_HILL_SCORING
			}
		}
	}
	return SCOUTING_SCORING
}

func (m *Map) FriendliesInRangeOf(p Point) []Point {
	reply := make([]Point, 0)
	for _, a := range m.myAnts {
		if a.CouldInfluence(p, m) {
			reply = append(reply, a.Point)
		}
	}
	if len(reply)>0 {
		m.enemyCombatants = append(m.enemyCombatants, p)
	}
	return reply
}

func (a Point ) CouldInfluence(b Point, m *Map) bool {
	isDry := func(p Point) bool {
		return !m.isWet(p)
	}
	aNext := filterPoints(a.NeighboursAndSelf(), isDry)
	bNext := filterPoints(b.NeighboursAndSelf(), isDry)

	for _, aa := range aNext {
		for _, bb := range bNext {
			if aa.CrowDistance2(bb) <= ATTACKRADIUS2 {
				return true
			}
		}
	}
	return false
}

var zoneNum int

func NewZone(e Point) *CombatZone {
	zoneNum += 1
	return &CombatZone{zone: zoneNum, enemy: []Point{e}}
}

func (cz *CombatZone) GroupCombat(m *Map) *GroupMove {

	if len(cz.friendly)+len(cz.enemy) > 7 {
		return cz.SimpleGroupCombat(m)
	}
	
	// For each of my possible moves, what could enemies do?

	bestMove := new(GroupMove)
	sh := cz.GetScoringHeuristic(m)
	for friendMove := range m.legalMoves(cz.friendly) {
		for enemyMove := range m.legalMoves(cz.enemy) {
			friendMove.score(enemyMove, sh)
		}
		bestMove.update(friendMove)
	}
	return bestMove
}

// There are too many ants in close proximity to try
// all combinations of moves.
// Simplify by:
//  * moving each of our ants independently
//  * assuming enemy ants stay still

func (cz *CombatZone) SimpleGroupCombat(m *Map) *GroupMove {
	occupied := occupiedMap(cz.friendly)
	dst := make([]Point, len(cz.friendly))
	var evalFn func(Point)int
	if len(cz.friendly) >= len(cz.enemy) {
		evalFn = func(p Point) int {
			// We outnumber them. Closer is better
			return - minDistance2(p, cz.enemy)
		}
	} else {
		// outnumbered: further away is better
		evalFn = func(p Point) int {
			return minDistance2(p, cz.enemy)
		}
	}
	
	for i, a := range cz.friendly {
		possibilities := nextMoves(a, m, occupied)
		var next Point
		if len(possibilities) == 0 {
			next = a
		} else {
			next = bestStep(possibilities, evalFn)
		}
		dst[i] = next
		occupied[a.loc()] = false
		occupied[next.loc()] = true
	}
	
	// FIXME: check expected K/D ratio, back out if this
	// move is crazy
	return &GroupMove{dst:dst}
}

func bestStep(alt []Point, evalFn func(Point)int) Point {
	bestScore := 0
	bestP := Point{}
	for i,p := range alt {
		score := evalFn(p)
		if i==0 || score > bestScore {
			bestP = p
			bestScore = score
		}
	}
	return bestP
}
func nextMoves(p Point, m *Map, occupied map[Location] bool)[]Point {
	start := p.NeighboursAndSelf()
	return filterPoints(start, func(p Point) bool {
		return !occupied[p.loc()] && !m.isWet(p)
	})
}

func occupiedMap(points []Point) map[Location]bool {
	reply := make(map[Location]bool)
	for _, p := range points {
		reply[p.loc()] = true
	}
	return reply
}


func (m *Map) legalMoves(orig []Point) chan GroupMove {
	ch := make(chan GroupMove)
	go func() {
		dst := make([]Point, len(orig))
		ps := NewPatternSet(orig)
		legal2(m, orig, dst, ps, 0, ch)
		close(ch)
	}()
	return ch
}

func legal2(m *Map, orig, dst []Point, patternSet *PatternSet, pos int, ch chan GroupMove) {
	src, orig := orig[0], orig[1:]
	allNeighbours := []Point{
		src,
		Point{src.r + 1, src.c},
		Point{src.r - 1, src.c},
		Point{src.r, src.c + 1},
		Point{src.r, src.c - 1},
	}
	// log.Printf("allNeighbours: %v", allNeighbours)
	for _, p := range allNeighbours {
		p.sanitise()
		if m.isWet(p) {
			continue
		}
		dst[pos] = p
		if patternSet.Seen(dst[:pos+1]) {
			continue
		}
		// log.Printf("dst: %v, pos:%v, p:%v", dst, pos, p)
		if len(orig) == 0 {
			dstCopy := make([]Point, len(dst))
			copy(dstCopy, dst)
			ch <- GroupMove{dst: dstCopy}
		} else {
			legal2(m, orig, dst, patternSet, pos+1, ch)
		}
	}
}

func (gm *GroupMove) update(om GroupMove) {
	gm.evaluated += om.evaluated
	if !gm.updated || om.worst > gm.worst || (om.worst == gm.worst && om.average > gm.average) {
		gm.worst = om.worst
		gm.average = om.average
		gm.updated = true
		gm.evaluated += om.evaluated
		gm.dst = om.dst
	}
}

func (gm *GroupMove) score(em GroupMove, sh ScoringHeuristic) {
	enemyFocus := make([]int, len(em.dst))
	friendlyFocus := make([]int, len(gm.dst))

	// Count focus
	for f, fp := range gm.dst {
		for e, ep := range em.dst {
			if fp.CrowDistance2(ep) <= ATTACKRADIUS2 {
				enemyFocus[e] += 1
				friendlyFocus[f] += 1
			}
		}
	}

	// Identify bodies
	enemyDead := make([]bool, len(em.dst))
	friendlyDead := make([]bool, len(gm.dst))

	for f, fp := range gm.dst {
		for e, ep := range em.dst {
			if fp.CrowDistance2(ep) <= ATTACKRADIUS2 {
				if enemyFocus[e] >= friendlyFocus[f] {
					enemyDead[e] = true
				}
				if enemyFocus[e] <= friendlyFocus[f] {
					friendlyDead[f] = true
				}
			}
		}
	}

	// Count bodies
	nEnemyDead, nFriendlyDead := countBool(enemyDead), countBool(friendlyDead)

	score := nEnemyDead*sh.deadEnemy + nFriendlyDead*sh.deadFriendly
	gm.evaluated += 1
	gm.total += score
	gm.average = float64(gm.total) / float64(gm.evaluated)
	if !gm.updated {
		gm.updated = true
		gm.worst = score
	} else {
		gm.worst = min(score, gm.worst)
	}
	//log.Printf("%s.score(%s) => %d", gm, &em, score)
}

func countBool(slice []bool) int {
	count := 0
	for _, b := range slice {
		if b {
			count += 1
		}
	}
	return count
}

