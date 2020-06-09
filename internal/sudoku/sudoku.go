package sudoku

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
)

// Puzzle represents a Sudoku puzzle.
//
// It is a 9x9 grid, stored as an array in row-major order.
type Puzzle [9 * 9]uint8

// String returns a pretty string representation of the puzzle.
func (p Puzzle) String() string {
	var b strings.Builder
	const sep = "+-------+-------+-------+"
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Fprintln(&b, sep)
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				fmt.Fprint(&b, "| ")
			}
			fmt.Fprintf(&b, "%d ", p[i*9+j])
		}
		fmt.Fprintln(&b, "|")
	}
	fmt.Fprint(&b, sep)
	return b.String()
}

// GoString implements fmt.GoStringer.
func (p Puzzle) GoString() string {
	var b bytes.Buffer
	b.WriteString("Puzzle{\n")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fprintf(&b, "%d,", p[i*9+j])
		}
		b.WriteByte('\n')
	}
	b.WriteByte('}')
	out, err := format.Source(b.Bytes())
	if err != nil {
		panic(err)
	}
	return string(out)
}

// Scan implements fmt.Scanner.
func (p *Puzzle) Scan(state fmt.ScanState, verb rune) error {
	for i := 0; i < len(p); i++ {
		r, _, err := state.ReadRune()
		if err != nil {
			return err
		}
		p[i] = uint8(r - '0')
		if p[i] > 9 {
			return fmt.Errorf("invalid input: %q", string(r))
		}
	}
	return nil
}

// Solve returns a solution to the puzzle. If no solution is possible, ok will
// be false.
func (p Puzzle) Solve() (s Puzzle, ok bool) {
	if !p.isValid() {
		return p, false
	}
	return p.solve()
}

func (p Puzzle) solve() (s Puzzle, ok bool) {
	i := p.firstEmptyIndex()
	if i == -1 {
		// p is complete
		return p, true
	}
	for _, n := range p.candidatesFor(i) {
		p[i] = n
		s, ok = p.solve()
		if ok {
			return
		}
	}
	return s, false
}

func (p Puzzle) isValid() bool {
	// check rows
	for i := 0; i < 9; i++ {
		var z uint16 // bitmap for numbers seen
		for j := 0; j < 9; j++ {
			n := p[9*i+j]
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
			n := p[9*i+j]
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
			var z uint16 // bitmap for numbers seen
			for i := a.start; i < a.end; i++ {
				for j := b.start; j < b.end; j++ {
					n := p[9*i+j]
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

func (p Puzzle) firstEmptyIndex() int {
	for i, n := range p {
		if n == 0 {
			return i
		}
	}
	return -1
}

func (p Puzzle) candidatesFor(i int) []uint8 {
	var z uint16 // bitmap
	i, j := i/9, i%9
	// visit row
	for x := 0; x < 9; x++ {
		n := p[9*i+x]
		z |= 1 << n
	}
	// visit col
	for x := 0; x < 9; x++ {
		n := p[9*x+j]
		z |= 1 << n
	}
	// visit square
	for x := i / 3 * 3; x < i/3*3+3; x++ {
		for y := j / 3 * 3; y < j/3*3+3; y++ {
			n := p[9*x+y]
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
