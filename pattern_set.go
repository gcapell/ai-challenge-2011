package main

type PatternSet struct {
	point2bit map[Location]uint
	seen map[uint]bool
}

func NewPatternSet (points []Point) *PatternSet {
	point2bit := make(map[Location]uint)
	var nextBit uint
	
	for _,p := range points {
		for _, p2 := range p.Neighbours(1) {
			loc := p2.loc()
			if _, ok := point2bit[loc]; !ok {
				point2bit[loc] = nextBit
				nextBit += 1
			}
		}
	}
	return &PatternSet{point2bit, make(map[uint]bool)}
}

// Map slice of points to bitmask
func (ps *PatternSet) pointsToMask (points [] Point) (mask uint) {
	for _, p := range points {
		b := ps.point2bit[p.loc()]
		mask |= 1<< b
	}
	return mask
}

// Have we already seen these points? 
// (If not, mark as seen)
func (ps *PatternSet) Seen (points [] Point ) bool {
	mask := ps.pointsToMask(points)
	if ps.seen[mask] {
		return true
	}
	ps.seen[mask] = true
	return false
}
