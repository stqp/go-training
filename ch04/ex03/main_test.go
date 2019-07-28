package main

import (
	"testing"
)

type Cases struct {
	arg      myarray
	expected myarray
}

func TestReverse(t *testing.T) {

	cases := []Cases{
		{arg: myarray{1, 2, 3, 4, 5, 6, 7, 8}, expected: myarray{8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, c := range cases {
		reverse(&c.arg)
		if !equals(&c.arg, &c.expected) {
			t.Error("Expect:", c.expected, "Actual:", c.arg)
		}
	}

}
