package inpt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var scanner *Scanner

func init() {
	scanner = NewScanner(os.Stdin)
}

// Normalize trims the leading and trailing whitespace of a string and makes all
// of its characters lowercase.
func Normalize(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return str
}

// Scanner is a wrapper around bufio.Scanner with convenient methods for
// collecting user input and cleaning it up.
type Scanner struct {
	*bufio.Scanner
	done bool
}

// NewScanner returns a Scanner of the supplied reader.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{Scanner: bufio.NewScanner(r)}
}

func (s *Scanner) Scan() bool {
	more := s.Scanner.Scan()
	s.done = !more
	return more
}

func (s *Scanner) More() bool {
	return !s.done
}

// Line reads a line from the scanner and trims the whitespace around it.
func (s *Scanner) Line() (string, error) {
	s.Scan()
	if err := s.Err(); err != nil {
		return "", fmt.Errorf("could not read line: %w", err)
	}
	result := s.Text()
	result = strings.TrimSpace(result)
	return result, nil
}

// LineInts reads a line of integers separated by commas
func (s *Scanner) IntLine() ([]int, error) {
	line, err := s.Line()
	if err != nil {
		return nil, err
	}
	intStrs := strings.Split(line, ",")
	var result []int
	for _, n := range intStrs {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

// Confirm reads a line and returns true if it is "y" or false if
// it is anything else.
func (s *Scanner) Confirm() (bool, error) {
	input, err := s.Line()
	if err != nil {
		return false, fmt.Errorf("could not get user confirmation: %w", err)
	}
	input = strings.ToLower(input)
	if input != "y" {
		return false, nil
	}
	return true, nil
}

// Line reads a line from stdin and trims the whitespace around it.
func Line() (string, error) {
	return scanner.Line()
}

// Confirm reads input from the user and returns true if it is "y" or false if
// it is anything else.
func Confirm() (bool, error) {
	return scanner.Confirm()
}
