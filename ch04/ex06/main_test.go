package main

import (
	"testing"
)

type Cases struct {
	arg      []string
	expected []string
}

func TestRemoveDuplicate(t *testing.T) {

	cases := []Cases{
		{arg: []string{"aa", "bb"}, expected: []string{"aa", "bb"}},
		{arg: []string{"aa", "aa", "bb"}, expected: []string{"aa", "bb"}},
		{arg: []string{"aa", "aa", "aa", "bb"}, expected: []string{"aa", "bb"}},
		{arg: []string{"aa", "bb", "aa", "bb"}, expected: []string{"aa", "bb", "aa", "bb"}},
	}

	for _, c := range cases {
		actual := removeDuplicate(c.arg)
		if !equals(actual, c.expected) {
			t.Error("Expect:", c.expected, "Actual:", actual)
		}
	}

}
