package sudoku

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

var (
	// testData has real puzzles and their respective solution.
	testData []Pair

	// testDataOnce is used to load testData only once.
	testDataOnce sync.Once
)

// A Pair is a puzzle and its solution.
type Pair struct {
	Puzzle   Puzzle
	Solution Puzzle
}

func (p Pair) String() string {
	var b strings.Builder
	b.WriteString("        Puzzle                     Solution")
	puzzleLines := strings.Split(p.Puzzle.String(), "\n")
	solutionLines := strings.Split(p.Solution.String(), "\n")
	for i := range puzzleLines {
		b.WriteByte('\n')
		b.WriteString(puzzleLines[i])
		b.WriteString("   ")
		b.WriteString(solutionLines[i])
	}
	return b.String()
}

// mustLoadTestData ensures that testData is loaded with test data or panics.
func mustLoadTestData() {
	testDataOnce.Do(func() {
		f, err := os.Open(filepath.Join("testdata", "test.csv"))
		if err != nil {
			panic(err)
		}
		var p, s Puzzle
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if scanner.Bytes()[0] == '#' {
				continue // skip comment lines
			}
			_, err := fmt.Fscanf(bytes.NewReader(scanner.Bytes()), "%v,%v", &p, &s)
			if err != nil {
				panic(err)
			}
			testData = append(testData, Pair{Puzzle: p, Solution: s})
		}
		if err := scanner.Err(); err != nil {
			panic(fmt.Errorf("reading %q: %s", f.Name(), err))
		}
		// sanity checks
		if len(testData) < 30 {
			panic(fmt.Errorf("found %d puzzles, want 30 or more", len(testData)))
		}
		for i, pair := range testData {
			if pair.Puzzle == pair.Solution {
				panic(fmt.Errorf("pair %d: puzzle equals solution:\n%s", i, pair.Puzzle))
			}
		}
		if testData[0] == testData[1] {
			panic("read repeated data")
		}
	})
}

func TestSolveTestData(t *testing.T) {
	mustLoadTestData()

	for _, p := range testData {
		t.Run("", func(t *testing.T) {
			got, ok := p.Puzzle.Solve()
			if !ok {
				t.Fatalf("solution not found, got:\n%s", got)
			}
			if got != p.Solution {
				t.Fatalf("got %v, want %v", got, p.Solution)
			}
		})
	}
}
