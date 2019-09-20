package main

import (
	"flag"
	"os"
	"testing"
)

type Case struct {
	arg      string
	expected string
}

func TestMain(t *testing.T) {

	var temp = CelsiusFlag("temp", 0.0, "")

	cases := []Case{
		{arg: "99C", expected: "99°C"},
		{arg: "99°C", expected: "99°C"},
		{arg: "99F", expected: "37.22222222222222°C"},
		{arg: "99°F", expected: "37.22222222222222°C"},
		{arg: "99K", expected: "-174.14999999999998°C"},
		{arg: "99°K", expected: "-174.14999999999998°C"},
	}

	for _, c := range cases {
		os.Args = []string{"", "-temp", c.arg}
		flag.Parse()
		if temp.String() != c.expected {
			t.Error("expected:", c.expected, "actual:", *temp)
		}
	}

}
