package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"foo/inpt"
)

//go:embed input.txt
var input string

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var sums = make([]int, 2000)

func sum(i int) int {
	if i == 0 {
		return 0
	}
	if sums[i] != 0 {
		return sums[i]
	}
	sums[i] = i + sum(i-1)
	return sums[i]
}

func checkFuel(crabs []int, pos int) int {
	total := 0
	for _, v := range crabs {
		travel := v - pos
		if travel < 0 {
			travel = -travel
		}
		total += sum(travel)
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
	answer := -1
	for i := range crabs {
		fuel := checkFuel(crabs, i)
		if fuel < answer || answer == -1 {
			answer = fuel
		}
	}
	fmt.Println("Minimum fuel for alignment:", answer)
}
