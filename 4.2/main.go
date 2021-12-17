package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"foo/inpt"
	"reflect"
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

const boardSize = 5

type board struct {
	numbers [][]int
	marked  [][]bool
}

func (b *board) mark(x int) {
	for i := range b.numbers {
		for j := range b.numbers[i] {
			if b.numbers[i][j] == x {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *board) hasWon() bool {
	for i := range b.numbers {
		lineCount := 0
		for j := range b.numbers[i] {
			if b.marked[i][j] == true {
				lineCount++
			}
		}
		if lineCount == boardSize {
			return true
		}
	}

	for i := range b.numbers {
		lineCount := 0
		for j := range b.numbers[i] {
			if b.marked[j][i] == true {
				lineCount++
			}
		}
		if lineCount == boardSize {
			return true
		}
	}

	return false
}

func (b *board) score(called int) int {
	sum := 0
	for i := range b.numbers {
		for j := range b.numbers[i] {
			if !b.marked[i][j] {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum * called
}

func newBoard() board {
	marked := make([][]bool, boardSize)
	for i := range marked {
		marked[i] = make([]bool, boardSize)
	}
	b := board{
		numbers: make([][]int, boardSize),
		marked:  marked,
	}
	return b
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

func removeBoard(boards []board, b board) []board {
	var result []board
	for i := range boards {
		if reflect.DeepEqual(boards[i], b) {
			continue
		}
		result = append(result, boards[i])
	}
	return result
}

func main() {
	data := bytes.NewBufferString(input)
	lines := inpt.NewScanner(data)
	// Get the Bingo numbers
	numLine, err := lines.Line()
	must(err)
	numStrs := strings.Split(numLine, ",")
	nums, err := convertToNums(numStrs)
	must(err)
	// Skip newline
	_, err = lines.Line()

	// Get the Bingo boards
	var boards []board
	for {
		b := newBoard()
		for i := 0; i < boardSize; i++ {
			line, err := lines.Line()
			b.numbers[i], err = convertToNums(strings.Fields(line))
			must(err)
		}
		boards = append(boards, b)
		// New Line
		_, err := lines.Line()
		must(err)
		if !lines.More() {
			break
		}
	}

	var lastWinner board
	var lastNum int
	for _, num := range nums {
		for _, board := range boards {
			board.mark(num)
			if board.hasWon() {
				lastWinner = board
				lastNum = num
				boards = removeBoard(boards, board)
			}
		}
	}
	fmt.Println("Final Score: ", lastWinner.score(lastNum))
}
