package main

import "testing"

type Cases struct {
	arg      []string
	expected string
}

func TestJoin(t *testing.T) {
	cases := []Cases{
		{arg: []string{"*", "a", "b"}, expected: "a*b"},
		{arg: []string{":", "a", "b", "c"}, expected: "a:b:c"},
		{arg: []string{"/", "aaa", "bbb", "ccc"}, expected: "aaa/bbb/ccc"},
	}

	for _, c := range cases {
		actual := join(c.arg[0], c.arg[1:]...)
		if actual != c.expected {
			t.Error("expected:", c.expected, "actual:", actual)
		}
	}
}
