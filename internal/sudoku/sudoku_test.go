package sudoku

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		name   string
		puzzle Puzzle
		valid  bool
	}{
		{
			"Empty",
			Puzzle{},
			true,
		},
		{
			"MagazineExample",
			Puzzle{
				0, 0, 0, 4, 0, 0, 3, 0, 0,
				0, 0, 0, 0, 6, 3, 0, 4, 9,
				4, 5, 3, 9, 0, 0, 0, 6, 7,
				0, 0, 8, 0, 0, 0, 0, 5, 0,
				0, 0, 0, 3, 0, 0, 0, 0, 0,
				6, 7, 0, 5, 9, 0, 0, 0, 2,
				0, 4, 0, 7, 0, 0, 2, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1, 3,
				9, 1, 0, 0, 3, 4, 0, 0, 0,
			},
			true,
		},
		{
			"InvalidRow",
			Puzzle{
				0, 0, 0, 4, 0, 0, 3, 0, 0,
				0, 0, 0, 0, 6, 3, 0, 4, 9,
				4, 5, 3, 9, 0, 0, 0, 6, 7,
				0, 0, 8, 0, 0, 0, 0, 5, 0,
				0, 0, 0, 3, 0, 0, 0, 0, 0,
				6, 7, 0, 5, 9, 0, 0, 0, 2,
				0, 4, 0, 7, 0, 0, 2, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1, 3,
				9, 1, 0, 0, 3, 4, 0, 0, 1, /*invalid*/
			},
			false,
		},
		{
			"InvalidCol",
			Puzzle{
				0, 0, 0, 4, 0, 0, 3, 0, 0,
				0, 0, 0, 0, 6, 3, 0, 4, 9,
				4, 5, 3, 9, 0, 0, 0, 6, 7,
				0, 0, 8, 0, 0, 0, 0, 5, 0,
				0, 0, 0, 3, 0, 0, 0, 0, 0,
				6, 7, 0, 5, 9, 0, 0, 0, 2,
				0, 4, 0, 7, 0, 0, 2, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1, 3,
				9, 1, 0, 0, 3, 4, 0, 0, 7, /*invalid*/
			},
			false,
		},
		{
			"InvalidSquare",
			Puzzle{
				0, 0, 0, 4, 0, 0, 3, 0, 0,
				0, 0, 0, 0, 6, 3, 0, 4, 9,
				4, 5, 3, 9, 0, 0, 0, 6, 7,
				0, 0, 8, 0, 0, 0, 0, 5, 0,
				0, 0, 0, 3, 5 /*invalid*/, 0, 0, 0, 0,
				6, 7, 0, 5, 9, 0, 0, 0, 2,
				0, 4, 0, 7, 0, 0, 2, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1, 3,
				9, 1, 0, 0, 3, 4, 0, 0, 0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := valid(tt.puzzle)
			want := tt.valid
			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestCandidatesFor(t *testing.T) {
	got := candidatesFor(Puzzle{}, 0)
	want := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSolve(t *testing.T) {
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
	got, ok := p.Solve()
	if !ok {
		t.Fatalf("solution not found, got:\n%s", got)
	}
	want := Puzzle{
		8, 6, 9, 4, 7, 5, 3, 2, 1,
		1, 2, 7, 8, 6, 3, 5, 4, 9,
		4, 5, 3, 9, 1, 2, 8, 6, 7,
		2, 3, 8, 1, 4, 7, 9, 5, 6,
		5, 9, 1, 3, 2, 6, 7, 8, 4,
		6, 7, 4, 5, 9, 8, 1, 3, 2,
		3, 4, 6, 7, 8, 1, 2, 9, 5,
		7, 8, 2, 6, 5, 9, 4, 1, 3,
		9, 1, 5, 2, 3, 4, 6, 7, 8,
	}
	if got != want {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
