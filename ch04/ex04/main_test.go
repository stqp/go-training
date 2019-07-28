package main

import (
	"testing"
)

type Cases struct {
	arg      []int
	expected []int
	by       int
}

func TestRotate(t *testing.T) {

	cases := []Cases{
		{arg: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: []int{6, 7, 8, 1, 2, 3, 4, 5}, by: 3},
		{arg: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: []int{7, 8, 1, 2, 3, 4, 5, 6}, by: 2},
		{arg: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: []int{8, 1, 2, 3, 4, 5, 6, 7}, by: 9},
	}
	for _, c := range cases {
		actual := RotateRight(c.arg, c.by)
		if !equals(actual, c.expected) {
			t.Error("Expect:", c.expected, "Actual:", actual)
		}
	}

	cases = []Cases{
		{arg: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: []int{4, 5, 6, 7, 8, 1, 2, 3}, by: 3},
		{arg: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: []int{3, 4, 5, 6, 7, 8, 1, 2}, by: 2},
		{arg: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: []int{2, 3, 4, 5, 6, 7, 8, 1}, by: 9},
	}
	for _, c := range cases {
		actual := RotateLeft(c.arg, c.by)
		if !equals(actual, c.expected) {
			t.Error("Expect:", c.expected, "Actual:", actual)
		}
	}

}
