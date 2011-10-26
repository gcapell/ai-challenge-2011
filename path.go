package main
import (
	"os"
	"container/heap"
	"container/vector"
)
type (
	Point struct {x,y int}

	Node struct {
		score int
		Point
	}

	myHeap struct {
		vector.Vector
	}
)

func (loc Location ) point () Point {
	iLoc := int(loc)
	return Point{iLoc / COLS, iLoc % COLS}
}

func (h *myHeap) Less(i, j int) bool { return h.At(i).(Node).score < h.At(j).(Node).score }

func (p Point) loc() Location {
	return toLoc(p.x, p.y)
}

func (m *Map) ShortestPath(srcLoc, dstLoc Location) ([] Location, os.Error) {
	src := srcLoc.point()

	h := &myHeap{}
	heap.Init(h)

	n := Node {3, src}
	h.Push(n)
	
	return []Location{1,2,4}, nil
}