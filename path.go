package main
import (
	"os"
	"container/heap"
	"container/vector"
	"fmt"
)
type (
	Node struct {
		estimate int
		length int
		p Point
	}

	myHeap struct {
		vector.Vector
	}
)

func (h *myHeap) Less(i, j int) bool { return h.At(i).(Node).estimate < h.At(j).(Node).estimate }

func unravelPath(back []Location, src, dst Point) []Point {
	// Follow breadcrumbs from 'dst' to 'src'
	reply := make([]Point, 0, ROWS + COLS)
	for p := dst; !p.Equals(src); p = back[p.loc()].point() {
		reply = append(reply, p)
	}
	return reply
}

func (m *Map) ShortestPath(srcLoc, dstLoc Location) ([] Point, os.Error) {
	src := srcLoc.point()
	dst := dstLoc.point()

	h := &myHeap{}
	heap.Init(h)

	heap.Push(h, Node{m.Distance(src, dst), 0, src})
	
	// Each entry points to previous point in path
	back := make([]Location, ROWS * COLS)

	for j :=0; j<len(back); j++ {
		back[j] = -1
	}

	for h.Len() != 0 {
		n := heap.Pop(h).(Node)
		pathLength := n.length + 1
		for _, p := range(m.DryNeighbours(n.p)) {
			// Have we already seen this point?
			if back[p.loc()] != -1 {
				continue
			}
			newNode := Node{m.Distance(p, dst) + pathLength, pathLength, p}
			back[p.loc()] = n.p.loc()
			if p.Equals(dst) {
				return unravelPath(back, src, dst), nil
			}
			heap.Push(h, newNode)
		}
	}
	return nil, fmt.Errorf("no path found")
}
