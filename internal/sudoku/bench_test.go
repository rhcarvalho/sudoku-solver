package sudoku

import (
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	p := Puzzle{
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
	for i := 0; i < b.N; i++ {
		p.Solve()
	}
}

func BenchmarkSolveTestDataEach(b *testing.B) {
	mustLoadTestData()

	for _, p := range testData {
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p.Puzzle.Solve()
			}
		})
	}
}

func BenchmarkSolveTestDataAll(b *testing.B) {
	mustLoadTestData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, p := range testData {
			p.Puzzle.Solve()
		}
	}
}
