package main

import (
	"os"
	"container/heap"
	"container/vector"
	"fmt"
)

const (
	EXPN_NONE = iota
	EXPN_HORIZONTAL
	EXPN_VERTICAL
)

type (
	Node2 struct {
		estimate float64
		length   int
		Point
		prev	*Node2	// backlinks to reconstruct path
		expanded
	}

	myHeap2 struct {
		vector.Vector
	}
)

func (h *myHeap) Less(i, j int) bool { return h.At(i).(Node2).estimate < h.At(j).(Node2).estimate }

func (src Point) ShortestPath2(dst Point, m *Map) ([]Point, os.Error) {

	h := &myHeap2{}
	heap.Init(h)

	heap.Push(h, Node{src.CrowDistance(dst), 0, src, nil})

	// Each entry points to previous point in path
	seen := make(map[Location]bool)

	for h.Len() != 0 {
		n := heap.Pop(h).(Node2)
		for n2 := range n.expand(m, dst) {
			if seen[n2.loc()] {
				continue
			}
			if n2.Equals(dst) {
				return n2.path()
			}
			heap.Push(h, newNode)
		}
		}
	}
	return nil, fmt.Errorf("no path found")
}

func (n *Node2) expand(m *Map, dst Point) chan *Node2 {
	c := make(chan *Node2)
	if n.expanded != EXPN_VERTICAL {
		go n.expandDir(m, dst, c, EAST)
		go n.exapndDir(m, dst, c, WEST)
	}
	if n.expanded != EXPN_HORIZONTAL {
		go n.expandDir(m, dst, c, NORTH)
		go n.expandDir(m, dst, c, SOUTH)
	}
	return c
	newNode := Node{p.CrowDistance(dst) + float64(pathLength), pathLength, p}
	
}

func (n *Node2) expandDir(m *Map, dst Point, c chan *Node2, d Direction) {
	p := n.Point
	for p:= n.Point; !m.isWet(p); p := d.East(p) {
		for _, d2 := range d.Orthogonal() {
		for p2 := d.North(p); !m.isWet(p); p2 := d.North(p2){
			if isCorner(p2, d, m, dst) {
				c <- n.Child(p2, dst, d)
			}
		}
	}
}
