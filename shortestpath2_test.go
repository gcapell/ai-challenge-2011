package main

import (
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"
	"testing"
	"fmt"
)

func (m *Map) waterPoints() []Point {
	points := make([]Point, 0, ROWS * COLS)
	for row := 0 ; row < ROWS; row ++ {
		for col := 0; col < COLS; col ++ {
			p := Point{row,col}
			if m.isWet(p) {
				points = append(points,p)
			}
		}
	}
	return points
}

func points2js(points []Point)string {
	s := make([]string,len(points))
	for j, p := range(points) {
		s[j] = fmt.Sprintf("[%d,%d]", p.r, p.c)
	}
	return strings.Join(s, ",")
}


func TestShortestPath2(t *testing.T) {
	m := readMap("../tools/maps/maze/maze_02p_01.map")
	log.Printf("map: \n%s", points2js(m.waterPoints()))
	src, dst := Point{0,20}, Point{20,61}
	path, err := src.ShortestPath(dst, m)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("path: %s -> [%s] -> %s", src, points2js(path), dst)
}

func readMap(filename string) *Map {
	fp, err := os.Open(filename)
	if err != nil {
		log.Panicf("%s opening %s", err, filename)
	}
	reader := bufio.NewReader(fp)
	data := make(map[string]int)
	var m *Map
	row := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fields := strings.Fields(string(line))
		if fields[0] == "m" {
			if m == nil {
				m = new (Map)
				m.Init(data["rows"], data["cols"], 25, false)
			}
			for col, c := range(fields[1]) {
				switch c {
				case '.', '0', '1':
					// nothing
				case '%':
					m.MarkWater(Point{row,col})
				default:
					log.Panicf("weird %s @ %d,%d", string(c), row, col)
				}
			}
			row += 1
		} else {
			data[fields[0]], _ = strconv.Atoi(fields[1])
		}
	}
	return m
}
