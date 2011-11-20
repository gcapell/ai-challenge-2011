package main

import (
	"os"
	"container/heap"
	"container/vector"
	"fmt"
)

type (
	Node struct {
		estimate float64
		length   int
		p        Point
	}

	myHeap struct {
		vector.Vector
	}
)

func (h *myHeap) Less(i, j int) bool { return h.At(i).(Node).estimate < h.At(j).(Node).estimate }

func reverse(points []Point) {
	for j, k := 0, len(points)-1; j < k; j, k = j+1, k-1 {
		points[k], points[j] = points[j], points[k]
	}
}

func unravelPath(back []Location, src, dst Point, pathLength int) []Point {
	// Follow breadcrumbs from 'dst' to 'src'
	reply := make([]Point, 0, pathLength)
	for p := dst; !p.Equals(src); p = back[p.loc()].point() {
		reply = append(reply, p)
	}
	reverse(reply)
	return reply
}

func (src Point) ShortestPath(dst Point, m *Map) ([]Point, os.Error) {

	h := &myHeap{}
	heap.Init(h)

	heap.Push(h, Node{src.CrowDistance(dst), 0, src})

	// Each entry points to previous point in path
	back := make([]Location, ROWS*COLS)

	INVALID_LOC := Location(ROWS * COLS + 1)

	for j := 0; j < len(back); j++ {
		back[j] = INVALID_LOC
	}

	for h.Len() != 0 {
		n := heap.Pop(h).(Node)
		pathLength := n.length + 1
		for _, p := range m.AccessibleNeighbours(n.p) {
			// Have we already seen this point?
			if back[p.loc()] != INVALID_LOC {
				continue
			}
			newNode := Node{p.CrowDistance(dst) + float64(pathLength), pathLength, p}
			back[p.loc()] = n.p.loc()
			if p.Equals(dst) {
				return unravelPath(back, src, dst, pathLength), nil
			}
			heap.Push(h, newNode)
		}
	}
	return nil, fmt.Errorf("no path found")
}
