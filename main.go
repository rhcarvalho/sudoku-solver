package main

import (
	"fmt"

	"github.com/rhcarvalho/sudoku-solver/internal/sudoku"
)

func main() {
	p := sudoku.Puzzle{}
	p.Solve()
	fmt.Println(p)

	p = sudoku.Puzzle{
		0, 0, 0, 4, 0, 0, 3, 0, 0,
		0, 0, 0, 0, 6, 3, 0, 4, 9,
		4, 5, 3, 9, 0, 0, 0, 6, 7,
		0, 0, 8, 0, 0, 0, 0, 5, 0,
		0, 0, 0, 3, 0, 0, 0, 0, 0,
		6, 7, 0, 5, 9, 0, 0, 0, 2,
		0, 4, 0, 7, 0, 0, 2, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 3,
		9, 1, 0, 0, 3, 4, 0, 0, 0,
	}
	p.Solve()
	fmt.Println(p)
}
