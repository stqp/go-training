package main

import "testing"

type Cases struct {
	arg      []int
	expected int
}

func TestMax(t *testing.T) {

	cases := []Cases{
		{arg: []int{}, expected: 0},
		{arg: []int{1, 2}, expected: 2},
		{arg: []int{1, 2, 3, 4, 5}, expected: 5},
		{arg: []int{100, -10, 999, 1, 6}, expected: 999},
	}

	for _, c := range cases {
		actual := max(c.arg...)
		if actual != c.expected {
			t.Error("expected:", c.expected, "actual:", actual)
		}
	}
}

func TestMin(t *testing.T) {

	cases := []Cases{
		{arg: []int{}, expected: 0},
		{arg: []int{1, 2}, expected: 1},
		{arg: []int{1, 2, 3, 4, 5}, expected: 1},
		{arg: []int{100, -10, 999, 1, 6}, expected: -10},
	}

	for _, c := range cases {
		actual := min(c.arg...)
		if actual != c.expected {
			t.Error("expected:", c.expected, "actual:", actual)
		}
	}
}
