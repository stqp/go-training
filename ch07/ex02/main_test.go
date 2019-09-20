package main

import (
	"os"
	"testing"
)

type Cases struct {
	arg      []byte
	expected int64
}

func TestCountWriter(t *testing.T) {

	cases := []Cases{
		{arg: []byte("12345"), expected: int64(5)},
		{arg: []byte("hello world"), expected: int64(11)},
	}
	expectedCount := int64(0)

	cw, count := CountingWriter(os.Stdout)
	for _, c := range cases {
		cw.Write(c.arg)
		expectedCount += c.expected
		if *count != expectedCount {
			t.Error("Expected:", expectedCount, "Actual:", *count)
		}
	}

}
