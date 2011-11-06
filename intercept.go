package main

import (
	"math"
)

func (defender Point) intercept(enemy, victim Point) Point {
	if defender.intercepts(victim, enemy) {
		return enemy
	}
	return between(victim, enemy)
}

func linear(a, p, b int) (ratio float64, between bool) {
	if abs(a-b) < 4 {
		return 0, abs(p-(a+b)/2) < 2
	}
	if sign(b-a) != sign(b-p) {
		return 0, false
	}
	if sign(b-a) != sign(p-a) {
		return 0, false
	}

	return float64(abs(b-p)) / float64(abs(b-a)), true
}

func similarRatio(r1, r2 float64) bool {
	return r1 == 0 || r2 == 0 || math.Fabs(r1-r2) < .05
}

// Is 'p' close enough to 'between' a and b?
func (p Point) intercepts(a, b Point) bool {
	if p.CrowDistance(a) < 4 || p.CrowDistance(b) < 4 {
		return true
	}
	rowRatio, rowBetween := linear(a.r, p.r, b.r)
	columnRatio, columnBetween := linear(a.c, p.c, b.c)

	return rowBetween && columnBetween && similarRatio(rowRatio, columnRatio)
}

// FIXME - eventually cope with warping
func between(a, b Point) Point {
	return Point{(a.r + b.r) / 2, (a.c + b.c) / 2}
}

