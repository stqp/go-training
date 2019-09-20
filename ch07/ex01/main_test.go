package main

import (
	"fmt"
	"testing"
)

type Case struct {
	arg      string
	expected string
}

func TestAdd(t *testing.T) {

	cases := []Case{
		{arg: "hello\nworld", expected: "{2 2}"},
		{arg: "aa  bb\n\ncc\ndd", expected: "{4 4}"},
	}

	var counter WordLineCounter

	for _, c := range cases {
		counter = WordLineCounter{}
		counter.Write([]byte(c.arg))
		actual := fmt.Sprint(counter)
		if actual != c.expected {
			t.Error("expected:", c.expected, "actual:", actual)
		}
	}

}
