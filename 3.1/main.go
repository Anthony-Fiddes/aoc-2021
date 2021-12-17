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
	totalLines := 0
	trueCounts := make([]int, 12)
	for {
		line, err := lines.Line()
		must(err)
		if !lines.More() {
			break
		}
		totalLines++

		for i, v := range line {
			if v == '1' {
				trueCounts[i]++
			}
		}
	}

	gammaStr := ""
	for _, v := range trueCounts {
		ones := v
		zeroes := totalLines - v
		if ones > zeroes {
			gammaStr += "1"
		} else {
			gammaStr += "0"
		}
	}
	gamma, err := strconv.ParseInt(gammaStr, 2, 16)
	must(err)

	epsilonStr := ""
	for _, v := range gammaStr {
		if v == '0' {
			epsilonStr += "1"
		} else {
			epsilonStr += "0"
		}
	}
	epsilon, err := strconv.ParseInt(epsilonStr, 2, 16)
	must(err)

	fmt.Println("Gamma:", gamma)
	fmt.Println("Epsilon:", epsilon)
	fmt.Println("Gamma x Epsilon:", gamma*epsilon)
}
