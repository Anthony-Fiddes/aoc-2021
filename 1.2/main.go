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

type window []*int

const windowLength = 3

func (w window) add(num *int) window {
	for i := range w {
		if w[i] == nil {
			w[i] = num
			return w
		}
	}
	w[0] = w[1]
	w[1] = w[2]
	w[2] = num
	return w
}

func (w window) sum() int {
	total := 0
	for _, v := range w {
		if v != nil {
			total += *v
		}
	}
	return total
}

func (w window) len() int {
	total := 0
	for _, v := range w {
		if v != nil {
			total += 1
		}
	}
	return total
}

func (w window) copy() window {
	other := newWindow()
	copy(other, w)
	return other
}

func newWindow() window {
	window := make(window, windowLength, windowLength)
	return window
}

func main() {
	data := bytes.NewBufferString(input)
	lines := inpt.NewScanner(data)
	total := 0
	prev := newWindow()
	current := newWindow()
	for {
		line, err := lines.Line()
		must(err)
		if !lines.More() {
			break
		}

		num, err := strconv.Atoi(line)
		must(err)
		current = current.add(&num)
		if current.len() < windowLength {
			continue
		}
		if prev.len() == windowLength && prev.sum() < current.sum() {
			total++
		}
		prev = current.copy()
	}
	fmt.Println("Total increases: ", total)
}
