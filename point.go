package main

type(
	//Location combines (Row, Col) coordinate pairs 
	// for use as keys in maps (and in a 1d array)
	Location int

	Point struct {r, c int}	// rows, columns
	Points []Point

)

func (s *Points) add(p Point) {
	*s = append(*s, p)
}

func (loc Location ) point () Point {
	iLoc := int(loc)
	return Point{ iLoc / COLS, iLoc % COLS}
}

func (p *Point) sanitise() {
	if p.r < 0 {
		p.r += ROWS
	}
	if p.r >= ROWS {
		p.r -= ROWS
	}
	if p.c < 0 {
		p.c += COLS
	}
	if p.c >= COLS {
		p.c -= COLS
	}
}

func (p Point) sanitised() Point {
	p.sanitise()
	return p
}

func (p Point) loc() Location {
	return Location(p.r * COLS + p.c)
}

func (p Point) Equals(r Point) bool {
	return p.r == r.r && p.c == r.c
}

func wrapDelta(a, b, wrap int) int {
	delta := a-b
	if delta<0 {
		delta = -delta
	}
	wrapped := wrap - delta
	// log.Printf("a: %d, b: %d, wrap: %d, delta: %d, wrapped: %d\n", a, b, wrap, delta, wrapped)
	if delta < wrapped {
		return delta
	}
	return wrapped
}

// Return (Manhattan) distance between two points,
// allowing for warping across edges
func (a Point) Distance( b Point) int {
	return wrapDelta(a.r, b.r, ROWS)+
		wrapDelta(a.c, b.c, COLS)
}

func (p Point) Neighbours(rad2 int) [] Point{
	reply := Points(make([]Point, rad2))
	if rad2 < 1 {
		return reply
	}
	for dr:=0; dr*dr<= rad2; dr++ {
		for dc := 0; dc*dc + dr  *dr <= rad2; dc++ {
			reply.add(Point{p.r + dr, p.c + dc}.sanitised())
			if(dr != 0) {
				reply.add(Point{p.r - dr, p.c + dc}.sanitised())
			}
			if (dc !=0) {
				reply.add(Point{p.r + dr, p.c - dc}.sanitised())
			}
			if (dr !=0 && dc != 0) {
				reply.add(Point{p.r - dr, p.c - dc}.sanitised())
			}
		}
	}
	return reply
}

