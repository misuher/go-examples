package merge_test

import (
	"testing"

	. "github.com/misuher/go-examples/algo-sort"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		array    []int
		expected []int
	}{
		{[]int{5, 2, 7, 1, 8}, []int{1, 2, 5, 7, 8}},
		{[]int{0, 1, 2}, []int{0, 1, 2}},
		{[]int{2, 1, 0}, []int{0, 1, 2}},
	}

	for _, elem := range tests {
		result := Merge(elem.array)
		if !intEq(result, elem.expected) {
			t.Errorf("expected %q but got %q", elem.expected, result)
		}
	}
}

func intEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for index, elem := range a {
		if elem != b[index] {
			return false
		}
	}
	return true
}
