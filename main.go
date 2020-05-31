package main

import (
	"fmt"
	"strings"
)

// State is a Sudoku game state. It consists of a 9x9 grid, stored row by row.
type State [9 * 9]uint8

func (s State) String() string {
	var b strings.Builder
	for i := 0; i < 9; i++ {
		fmt.Fprintf(&b, "%v\n", s[i*9:i*9+9])
	}
	return b.String()
}

func solve(s State) (r State, ok bool) {
	if !valid(s) {
		return s, false
	}
	if isComplete(s) {
		return s, true
	}
	i := firstEmptyIndex(s)
	for _, n := range candidatesFor(s, i) {
		s[i] = n
		r, ok = solve(s)
		if ok {
			return
		}
	}
	return State{}, false
}

func valid(s State) bool {
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

func isComplete(s State) bool {
	// assume that s is valid
	for _, n := range s {
		if n == 0 {
			return false
		}
	}
	return true
}

func firstEmptyIndex(s State) int {
	for i, n := range s {
		if n == 0 {
			return i
		}
	}
	return -1
}

func candidatesFor(s State, i int) []uint8 {
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

func main() {
	fmt.Println(solve(State{
		0, 0, 0, 4, 0, 0, 3, 0, 0,
		0, 0, 0, 0, 6, 3, 0, 4, 9,
		4, 5, 3, 9, 0, 0, 0, 6, 7,
		0, 0, 8, 0, 0, 0, 0, 5, 0,
		0, 0, 0, 3, 0, 0, 0, 0, 0,
		6, 7, 0, 5, 9, 0, 0, 0, 2,
		0, 4, 0, 7, 0, 0, 2, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 3,
		9, 1, 0, 0, 3, 4, 0, 0, 0,
	}))
}
