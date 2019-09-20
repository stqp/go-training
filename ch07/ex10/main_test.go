package main

import (
	"testing"
)

type Strategy []rune

func (x Strategy) Less(i, j int) bool {
	return x[i] < x[j]
}
func (x Strategy) Len() int {
	return len(x)
}
func (x Strategy) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func TestIsPalindrome(t *testing.T) {

	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "abcde", expected: false},
		{arg: "abcba", expected: true},
		{arg: "abcd", expected: false},
		{arg: "abba", expected: true},
		{arg: "こんにちは", expected: false},
		{arg: "しんぶんし", expected: true},
		{arg: "一扁桃石蓴蓴石桃扁一", expected: true},
	}

	for _, c := range cases {
		s := Strategy(c.arg[:])
		if IsPalindrome(s) != c.expected {
			t.Error()
		}
	}
}
