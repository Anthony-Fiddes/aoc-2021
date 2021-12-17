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

type position struct {
	x, y, aim int
}

func main() {
	data := bytes.NewBufferString(input)
	lines := inpt.NewScanner(data)
	var pos position
	for {
		line, err := lines.Line()
		must(err)
		if !lines.More() {
			break
		}

		fields := strings.Fields(line)
		direction, numStr := fields[0], fields[1]
		num, err := strconv.Atoi(numStr)
		must(err)
		switch direction {
		case "forward":
			pos.x += num
			pos.y += pos.aim * num
		case "down":
			pos.aim += num
		case "up":
			pos.aim -= num
		}
	}
	fmt.Println("Depth:", pos.y)
	fmt.Println("Horizontal Position:", pos.x)
	fmt.Println("Depth x Horizontal Position: ", pos.x*pos.y)
}
