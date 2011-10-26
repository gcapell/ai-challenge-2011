package main
import (
	"os"
)
type Point struct {x,y int}

func (p Point) loc() Location {
	return toLoc(p.x, p.y)
}

type Node struct {
	score int
	Point
}

func (a *Node) Less (b *Node) bool {
	return a.score < b.score
}

type NodeVector []Node

func (m *Map) ShortestPath(src, dst Location) ([] Location, os.Error) {
	return []Location{1,2,4}, nil
}