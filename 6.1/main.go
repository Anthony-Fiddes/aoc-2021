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

type population []int

func (p population) remove(i int) population {
	end := len(p) - 1
	p[i] = p[end]
	p = p[:end]
	return p
}

func (p population) next() population {
	var toAdd population
	for i := range p {
		if p[i] == 0 {
			toAdd = append(toAdd, 8)
			p[i] = 6
			continue
		}
		p[i]--
	}
	p = append(p, toAdd...)
	return p
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
	pop := population(nums)
	const numDays = 80
	for i := 0; i < numDays; i++ {
		pop = pop.next()
	}
	fmt.Println("Population after", numDays, ":", len(pop))
}
