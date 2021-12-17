package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"foo/inpt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type point struct {
	x, y int
}

type line struct {
	start, end point
}

func (l line) isHorizontal() bool {
	if l.start.y == l.end.y {
		return true
	}
	return false
}

func (l line) isVertical() bool {
	if l.start.x == l.end.x {
		return true
	}
	return false
}

const gridSize = 1000

type grid [][]int

func (g grid) addLine(l line) grid {
	if l.isHorizontal() {
		currX := l.start.x
		endX := l.end.x
		for {
			g[currX][l.start.y]++
			if currX < endX {
				currX++
			} else if currX == endX {
				break
			} else {
				currX--
			}
		}
	} else if l.isVertical() {
		currY := l.start.y
		endY := l.end.y
		for {
			g[l.start.x][currY]++
			if currY < endY {
				currY++
			} else if currY == endY {
				break
			} else {
				currY--
			}
		}
	} else {
		// It's diagonal
		curr := l.start
		end := l.end
		for {
			g[curr.x][curr.y]++
			if curr == end {
				break
			}
			if curr.y < end.y {
				curr.y++
			} else {
				curr.y--
			}
			if curr.x < end.x {
				curr.x++
			} else {
				curr.x--
			}
		}
	}
	return g
}

func newGrid() grid {
	result := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		result[i] = make([]int, gridSize)
	}
	return grid(result)
}

func parsePoint(s string) point {
	coords := strings.Split(s, ",")
	for i := range coords {
		coords[i] = strings.TrimSpace(coords[i])
	}
	result := point{}
	var err error
	result.x, err = strconv.Atoi(coords[0])
	must(err)
	result.y, err = strconv.Atoi(coords[1])
	must(err)
	return result
}

func parseLine(s string) line {
	points := strings.Split(s, "->")
	result := line{}
	result.start = parsePoint(points[0])
	result.end = parsePoint(points[1])
	return result
}

func main() {
	data := bytes.NewBufferString(input)
	input := inpt.NewScanner(data)
	g := newGrid()
	for {
		next, err := input.Line()
		must(err)
		if !input.More() {
			break
		}
		next = strings.TrimSpace(next)
		l := parseLine(next)
		fmt.Println(l)
		g = g.addLine(l)
	}
	total := 0
	for x := range g {
		for y := range g[x] {
			if g[x][y] > 1 {
				total++
			}
		}
	}
	fmt.Println("Final Score:", total)
}
