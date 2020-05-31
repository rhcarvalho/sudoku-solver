package sudoku

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	if !valid(Puzzle{}) {
		t.Fatalf("empty state should be valid")
	}
}

func TestCandidatesFor(t *testing.T) {
	got := candidatesFor(Puzzle{}, 0)
	want := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
