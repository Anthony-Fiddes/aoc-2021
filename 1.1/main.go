package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"foo/inpt"
	"strconv"
)

//go:embed input.txt
var input string

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data := bytes.NewBufferString(input)
	lines := inpt.NewScanner(data)
	total := 0
	var prev *int
	for {
		line, err := lines.Line()
		must(err)
		if !lines.More() {
			break
		}

		num, err := strconv.Atoi(line)
		must(err)
		if prev != nil {
			if *prev < num {
				total++
			}
		}
		prev = &num
	}
	fmt.Println("Total increases: ", total)
}
