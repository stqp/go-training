package main

import (
	"testing"
)

func yesNoToBool(s string) bool {
	if s == "Yes" {
		return true
	}
	if s == "No" {
		return false
	}
	return false
}
func TestMain(t *testing.T) {

	cases := [][]string{
		{"abc", "cba", "Yes"},
		{"anagrams", "ARS MAGNA", "Yes"},
		{"Statue of Liberty", "built to stay free", "Yes"},
		{"Christmas", "trims cash", "Yes"},
		{"narcissism", "man's crisis", "Yes"},
		{"abc", "a b c", "No"},
		{"abc", "a'''bc", "No"},
		{"abc", "def", "No"},
	}
	for _, c := range cases {
		s1 := c[0]
		s2 := c[1]
		expect := yesNoToBool(c[2])
		if isAnagram(s1, s2) != expect {
			t.Error("fail", expect, isAnagram(s1, s2))
		}
	}

}
