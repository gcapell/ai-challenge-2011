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

func TestShortestPath2(t *testing.T) {
	m := readMap("../tools/maps/maze/maze_02p_01.map")
	src, dst := Point{15,23}, Point{33,58}
	path, err := src.ShortestPath(dst, m)
	if err != nil {
		t.Fatal(err)
	}

	fp, err := os.Create("data.js")

	if err != nil {
		log.Panic(err, "opening data.js")
	}
	path2, err := src.ShortestPath2(dst, m, fp)
	if err != nil {
		t.Fatal(err)
	}
	data := map[string][]Point { "water": m.waterPoints(), "path": path, "path2": path2}
	for k,v := range data {
		fmt.Fprintf(fp, "%s = %s;\n\n", k, points2js(v))
	}
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
