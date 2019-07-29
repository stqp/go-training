package main

import (
	"testing"
)

type Cases struct {
	arg      []rune
	expected []rune
}

func TestRemoveDuplicate(t *testing.T) {

	cases := []Cases{
		{arg: []rune("　こんにちは 世界 "), expected: []rune(" 界世 はちにんこ　")},
		{arg: []rune("おはよう"), expected: []rune("うよはお")},
	}

	for _, c := range cases {
		encodedRunes := encodeRunes(c.arg)
		replaced := reverseRunes(encodedRunes)
		actual := []rune(string(replaced))
		if !equals(actual, c.expected) {
			t.Error("Expect:", c.expected, "Actual:", actual)
		}
	}

}
