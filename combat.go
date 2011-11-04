package main

import (
	"log"
	"fmt"
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
	for _, cz := range m.FindCombatZones() {
		if m.deadlineExpired() {
			break
		}
		bestMove := cz.GroupCombat(m)
		if bestMove != nil {
			cz.MakeMove(m, bestMove)
		}
	}
}

func (cz *CombatZone) MakeMove(m *Map, bestMove *GroupMove) {
	// log.Printf("groupCombat friends: %v, enemies: %v, best: %v", cz.friendly, cz.enemy, bestMove)
	for i, p := range cz.friendly {
		dst := bestMove.dst[i]
		ant := m.myAnts[p.loc()]
		ant.moveToPoint(m, dst, "combat")
	}
}

func (m *Map) FindCombatZones() []*CombatZone {
	zoneNum = 0
	zones := make([]*CombatZone, 0)

	for _, e := range m.enemies {
		myZone := NewZone(e)
		merging := false
		for _, a := range m.FriendliesInRangeOf(e) {
			found := false
			// Is this friendly part of an existing zone?
			for _, z := range zones {
				if a.In(z.friendly) {
					found = true
					if !merging {
						merging = true
						z.enemy = append(z.enemy, e)
						z.friendly = append(z.friendly, myZone.friendly...)
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
			if p.CrowDistance2(h) <= VIEWRADIUS2 {
				return NEAR_OUR_HILL_SCORING
			}
		}
	}
	return SCOUTING_SCORING
}

func (m *Map) FriendliesInRangeOf(p Point) []Point {
	reply := make([]Point, 0)
	for _, a := range m.myAnts {
		if p.CouldInfluence(a.p) {
			reply = append(reply, a.p)
		}
	}
	return reply
}

func (a Point) CouldInfluence(b Point) bool {
	dr, dc := a.Distance(b)
	// Move two steps closer
	for j := 0; j < 2; j++ {
		if dr > dc {
			dr -= 1
		} else {
			dc -= 1
		}
	}
	return dr*dr+dc*dc <= ATTACKRADIUS2
}

var zoneNum int

func NewZone(e Point) *CombatZone {
	zoneNum += 1
	return &CombatZone{zone: zoneNum, enemy: []Point{e}}
}

func (cz *CombatZone) GroupCombat(m *Map) *GroupMove {

	if len(cz.friendly)+len(cz.enemy) > 7 {
		log.Printf("group combat too hard, giving up")
		return nil
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
	// log.Printf("%s.score(%s) => %d", gm, &em, score)
}
