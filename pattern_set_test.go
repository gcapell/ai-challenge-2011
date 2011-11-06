package main

import (
	"testing"
)

func TestPatternSet(t *testing.T) {
	points := [] Point { {1,2}, {4,5}, {8,9}}
	
	ps := NewPatternSet(points)
	
	pattern := [] Point { {1,2}, {4,4}}
	if ps.Seen(pattern) {
		t.Fatal("Should not have seen %v yet", pattern)
	}
	
	if !ps.Seen(pattern) {
		t.Fatal("Now we should have seen %v", pattern)
	}
}
