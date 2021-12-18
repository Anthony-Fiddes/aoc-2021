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

// Looking back, I definitely could have done this with an array
type population map[int]int

func (p population) add(fish []int) {
	for _, v := range fish {
		p[v]++
	}
}

func (p population) next() population {
	next := newPopulation()
	for k, v := range p {
		if k == 0 {
			next[8] += v
			next[6] += v
			continue
		}
		next[k-1] += v
	}
	return next
}

func (p population) total() int {
	total := 0
	for _, v := range p {
		total += v
	}
	return total
}

func newPopulation() population {
	result := make(population, 9)
	return result
}

func convertToNums(nums []string) ([]int, error) {
	var result []int
	for _, n := range nums {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

func main() {
	data := bytes.NewBufferString(input)
	lines := inpt.NewScanner(data)
	initialState, err := lines.Line()
	must(err)
	numStrs := strings.Split(initialState, ",")
	nums, err := convertToNums(numStrs)
	must(err)
	pop := newPopulation()
	pop.add(nums)
	const numDays = 256
	for i := 0; i < numDays; i++ {
		pop = pop.next()
	}
	fmt.Println("Population after", numDays, ":", pop.total())
}
