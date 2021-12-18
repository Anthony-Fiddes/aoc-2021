package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"foo/inpt"
	"log"
	"sort"
)

//go:embed input.txt
var input string

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func checkFuel(crabs []int, pos int) int {
	total := 0
	for _, v := range crabs {
		used := v - pos
		if used < 0 {
			used = -used
		}
		total += used
	}
	return total
}

func min(nums ...int) int {
	m := nums[0]
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}

func main() {
	data := bytes.NewBufferString(input)
	input := inpt.NewScanner(data)
	crabs, err := input.IntLine()
	must(err)
	sort.Ints(crabs)
	middle := len(crabs) / 2
	if len(crabs) < 3 {
		log.Fatal("Input must have more than 3 elements")
	}
	less := checkFuel(crabs, crabs[middle-1])
	median := checkFuel(crabs, crabs[middle])
	greater := checkFuel(crabs, crabs[middle+1])
	answer := min(less, median, greater)
	fmt.Println("Minimum fuel for alignment:", answer)
}
