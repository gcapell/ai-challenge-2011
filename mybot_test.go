package main

import (
	"testing"
)

type Pairing struct{ p, q Point }

func TestAssign1(t *testing.T) {

	ROWS = 20
	COLS = 20

	assign1Helper(t,
		[]Point{{3, 4}, {9, 5}},
		[]Point{{1, 2}, {7, 8}, {4, 4}},
		[]Pairing{
			{Point{3, 4}, Point{4, 4}},
			{Point{9, 5}, Point{7, 8}},
		})
}

func assign1Helper(t *testing.T, antLocations []Point, targets []Point, expected []Pairing) {
	ants := make([]*Ant, len(antLocations))
	for i, p := range antLocations {
		ants[i] = &Ant{Point: p}
	}
	reply := assign1(ants, targets)
	if len(reply) != len(expected) {
		t.Errorf("Got %s, expected %v\n", reply, expected)
	}
	for i, r := range reply {
		e := expected[i]
		if !(e.p.Equals(r.ant.Point) && e.q.Equals(r.p)) {
			t.Errorf("%v != %v in %s\n", r, e, reply)
		}
	}

}
