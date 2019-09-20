package main

import (
	"testing"
)

type Case struct {
	arg      []int
	expected string
}

func TestString(t *testing.T) {

	cases := []Case{
		{arg: []int{1, 2, 3, 4, 5}, expected: "{1 2 3 4 5}"},
		{arg: []int{10, 1, 9, 2, 8, 3}, expected: "{1 2 3 8 9 10}"},
	}

	for _, c := range cases {
		var tre *tree
		for _, v := range c.arg {
			tre = add(tre, v)
		}
		if tre.String() != c.expected {
			t.Error("expected:", c.expected, "actual:", tre.String())
		}
	}

}
