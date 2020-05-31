package sudoku

import (
	"fmt"
	"strings"
)

// Puzzle represents a Sudoku puzzle.
//
// It is a 9x9 grid, stored as an array in row-major order.
type Puzzle [9 * 9]uint8

func (s Puzzle) String() string {
	var b strings.Builder
	for i := 0; i < 9; i++ {
		fmt.Fprintf(&b, "%v\n", s[i*9:i*9+9])
	}
	return b.String()
}

// Solve returns a solution to the puzzle. If no solution is possible, ok will
// be false.
func (s Puzzle) Solve() (r Puzzle, ok bool) {
	if !valid(s) {
		return s, false
	}
	if isComplete(s) {
		return s, true
	}
	i := firstEmptyIndex(s)
	for _, n := range candidatesFor(s, i) {
		s[i] = n
		r, ok = s.Solve()
		if ok {
			return
		}
	}
	return Puzzle{}, false
}

func valid(s Puzzle) bool {
	// check rows
	for i := 0; i < 9; i++ {
		var z uint16 // bitmap for numbers seen
		for j := 0; j < 9; j++ {
			n := s[9*i+j]
			if n == 0 {
				continue
			}
			if n > 9 {
				return false // out of range
			}
			if z&(1<<n) != 0 {
				return false // repeated number
			}
			z |= 1 << n
		}
	}
	// check cols
	for j := 0; j < 9; j++ {
		var z uint16 // bitmap for numbers seen
		for i := 0; i < 9; i++ {
			n := s[9*i+j]
			if n == 0 {
				continue
			}
			if n > 9 {
				return false // out of range
			}
			if z&(1<<n) != 0 {
				return false // repeated number
			}
			z |= 1 << n
		}
	}
	// check squares
	t := [...]struct{ start, end int }{
		{0, 3}, {3, 6}, {6, 9},
	}
	for _, a := range t {
		for _, b := range t {
			for i := a.start; i < a.end; i++ {
				var z uint16 // bitmap for numbers seen
				for j := b.start; j < b.end; j++ {
					n := s[9*i+j]
					if n == 0 {
						continue
					}
					if n > 9 {
						return false // out of range
					}
					if z&(1<<n) != 0 {
						return false // repeated number
					}
					z |= 1 << n
				}
			}
		}
	}
	return true
}

func isComplete(s Puzzle) bool {
	// assume that s is valid
	for _, n := range s {
		if n == 0 {
			return false
		}
	}
	return true
}

func firstEmptyIndex(s Puzzle) int {
	for i, n := range s {
		if n == 0 {
			return i
		}
	}
	return -1
}

func candidatesFor(s Puzzle, i int) []uint8 {
	var z uint16 // bitmap
	i, j := i/9, i%9
	// visit row
	for x := 0; x < 9; x++ {
		n := s[9*i+x]
		z |= 1 << n
	}
	// visit col
	for x := 0; x < 9; x++ {
		n := s[9*x+j]
		z |= 1 << n
	}
	// visit square
	for x := i / 3 * 3; x < i/3*3+3; x++ {
		for y := j / 3 * 3; y < j/3*3+3; y++ {
			n := s[9*x+y]
			z |= 1 << n
		}
	}
	var r []uint8
	for n := 1; n <= 9; n++ {
		if z&(1<<n) == 0 {
			r = append(r, uint8(n))
		}
	}
	return r
}
