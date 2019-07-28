package main

import (
	"testing"
)

type Cases struct {
	arg1   [32]byte
	arg2   [32]byte
	expect int
}

func TestPopCount(t *testing.T) {

	var ffbytes [32]byte
	for i := 0; i < 32; i++ {
		ffbytes[i] = 0xFF
	}
	cases := []Cases{
		{arg1: [32]byte{0x01}, arg2: [32]byte{0x02}, expect: 2},
		{arg1: [32]byte{0x01, 0xFF}, arg2: [32]byte{0x02}, expect: 10},
		{arg1: ffbytes, arg2: [32]byte{0x00}, expect: 256},
	}

	for _, c := range cases {
		actual := XorPopCount(c.arg1, c.arg2)
		if actual != c.expect {
			t.Error("Expect:", c.expect, "Actual:", actual)
		}
	}

}
