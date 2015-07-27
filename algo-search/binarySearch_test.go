package binarySearch_test

import (
	"testing"

	. "github.com/misuher/go-examples/algo"
)

func TestBsearch(t *testing.T) {
	var tests = []struct {
		sortedArray []int
		elemToFind  int
		expected    int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 4, 3},
		{[]int{0, 2, 4, 8, 10, 12, 14, 16}, 16, 7},
		{[]int{0, 2, 4, 8, 10, 12, 14, 16}, 0, 0},
		{[]int{0, 2, 4, 8, 10, 12, 14, 16}, 2, 1},
	}

	for _, elem := range tests {
		result := Bsearch(elem.sortedArray, elem.elemToFind)
		if result != elem.expected {
			t.Errorf("expected %q but got %q", elem.expected, result)
		}
	}
}
