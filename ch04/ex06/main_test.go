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
		{arg: []rune("こんにちは　世界"), expected: []rune("こんにちは 世界")},
		{arg: []rune("　こ　ん　に　ち　は　"), expected: []rune(" こ ん に ち は ")},
	}

	for _, c := range cases {
		encodedRunes := encodeRunes(c.arg)
		replaced := replaceUnicodeSpace(encodedRunes)
		actual := []rune(string(replaced))
		if !equals(actual, c.expected) {
			t.Error("Expect:", c.expected, "Actual:", actual)
		}
	}

}
