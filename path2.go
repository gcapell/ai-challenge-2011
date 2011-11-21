package main

import (
	"os"
	"container/heap"
	"container/vector"
	"fmt"
	"log"
	"strings"
	"io"
)

const (
	EXPN_NONE = iota
	EXPN_HORIZONTAL
	EXPN_VERTICAL
)

type (
	Delta struct {
		dr, dc	int	// forward
	}
	Node2 struct {
		estimate float64
		length   int
		Point
		prev	*Node2	// backlinks to reconstruct path
		direction	Direction	// direction to expand next
	}

	myHeap2 struct {
		vector.Vector
	}
)

var (
	DIRECTIONS = []Delta {
		{-1,0}, {0,1}, {1,0}, {0, -1},
	}
)

type Direction int
const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
	NODIRECTION
)

var directionNames = []string {"N", "E", "S", "W", "ND"}

func (d Direction) String()string {
	return directionNames[d]
}

func (d Direction) Next(p Point) Point {
	delta := DIRECTIONS[d]
	p.r += delta.dr
	p.c += delta.dc
	p.sanitise()
	return p
}

func (n*Node2)String() string {
	return fmt.Sprintf("(%d,%d) ->%s %.1f %d", n.r, n.c, n.direction, n.estimate, n.length)
}

func (n *Node2) hash() int {
	return int(n.direction) * (ROWS * COLS) + n.r * COLS + n.c
}

func (d Direction) Orthogonal() []Direction {
	return []Direction { d.Left(), d.Right()}
}

func (d Direction) Opposite() Direction {
	return (d+2)%4
}

func (d Direction) Left() Direction {
	return (d+3)%4
}

func (d Direction) Right() Direction {
	return (d+1)%4
}

func isCorner(p Point, m *Map, left, back Direction) bool {
	leftPoint := left.Next(p)
	if !m.isWet(leftPoint) {
		backLeftPoint := back.Next(leftPoint)
		if m.isWet(backLeftPoint) {
			return true
		}
	}
	return false
}

func (h *myHeap2) Less(i, j int) bool { return h.At(i).(*Node2).estimate < h.At(j).(*Node2).estimate }

func nodes2js(nodes []*Node2) string {
	bits := make([]string, len(nodes))
	for j, n := range nodes {
		var pr, pc int
		if n.prev != nil {
			pr, pc = n.prev.Point.r, n.prev.Point.c
		} else {
			pr, pc = 0, 0
		}
		bits[j] = fmt.Sprintf("[%d,%d,%d,%.1f,'%s', %d, %d]", n.r, n.c, n.length, n.estimate, n.direction, pr, pc)
	}
	return fmt.Sprintf("[ %s ]", strings.Join(bits, ","))
}

func (src Point) ShortestPath2(dst Point, m *Map, fp io.Writer) ([]Point, os.Error) {

	h := &myHeap2{}
	heap.Init(h)

	heap.Push(h, NewNode(src,dst))

	// Each entry points to previous point in path
	seen := make(map[Location]bool)

	expansions := make([]Point, 0)
	popped := make([]*Node2, 0)
	defer func() {
		fmt.Fprintf(fp, "popped = %s;\n\n", nodes2js(popped))
		fmt.Fprintf(fp, "expansions= %s;\n\n", points2js(expansions))
	}()
	
	for h.Len() != 0 {
		n := heap.Pop(h).(*Node2)
		popped = append(popped, n)

		for n2 := range n.expand(m, dst, seen) {
			log.Printf("%s -> %s", n, n2)
			expansions = append(expansions, n2.Point)
			if n2.Equals(dst) {
				return n2.path(), nil
			}
			heap.Push(h, n2)
		}
	}
	return nil, fmt.Errorf("no path found")
}

func (n *Node2) expand(m *Map, dst Point, seen map[Location] bool) chan *Node2 {
	c := make(chan *Node2)
	go func() {
		if n.direction == NODIRECTION {
			n.expandDir(c, m, dst, NORTH, seen)
			n.expandDir(c, m, dst, SOUTH, seen)
			n.expandDir(c, m, dst, EAST, seen)
			n.expandDir(c, m, dst, WEST, seen)
		} else {
			n.expandDir(c, m, dst, n.direction, seen)
		}
		close(c)
	}()
	return c
}

func points2js(points []Point)string {
	s := make([]string,len(points))
	for j, p := range(points) {
		s[j] = fmt.Sprintf("[%d,%d]", p.r, p.c)
	}
	return fmt.Sprintf("[ %s ]", strings.Join(s, ","))
}


func (n *Node2) expandDir(c chan *Node2, m *Map, dst Point, d Direction, seen map[Location]bool) {
	log.Printf("expanding %s direction %s", n, d)
	for p := n.Point; !m.isWet(p); p = d.Next(p) {
		for _, d2 := range d.Orthogonal() {
			back := d2.Opposite()
			left := d2.Left()
			right := d2.Right()

			for p2 := d2.Next(p); !m.isWet(p2) && !p2.Equals(p); p2 = d2.Next(p2){
				if seen[p2.loc()] {
					continue
				} else {
					seen[p2.loc()] = true
				}
				if p2.Equals(dst) {
					c <- n.Child(p2, dst, d)
					return
				}
				if isCorner(p2, m, left, back) {
					c <- n.Child(p2, dst, left)
				}
				if isCorner(p2, m, right, back) {
					c <- n.Child(p2, dst, right)
				}
			}
		}
	}
}

func (n *Node2) path() []Point {
	reply := make([]Point,0)
	for n != nil {
		reply = append(reply, n.Point)
		n = n.prev
	}
	return reply
}

func NewNode(src, dst Point) *Node2 {
	return &Node2{src.CrowDistance(dst), 0, src, nil, NODIRECTION}
}

func (n *Node2) Child(p, dst Point, d Direction) *Node2 {
	length := n.length + n.ManhattanDistance(p)
	estimate := float64(length) + p.CrowDistance(dst)
	return &Node2{ estimate,length, p, n, d}
}
