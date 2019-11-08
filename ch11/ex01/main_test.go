package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input  string
		counts map[rune]int
		utflen [utf8.UTFMax + 1]int
	}{
		{
			"abcdab",
			map[rune]int{
				'a': 2,
				'b': 2,
				'c': 1,
				'd': 1,
			},
			[utf8.UTFMax + 1]int{0, 6, 0, 0, 0},
		},
		{
			"abcbaあああいいいい",
			map[rune]int{
				'a': 2,
				'b': 2,
				'c': 1,
				'あ': 3,
				'い': 4,
			},
			[utf8.UTFMax + 1]int{0, 5, 0, 7, 0},
		},
	}
	for _, test := range tests {
		in := strings.NewReader(test.input)
		if c, u := Charcount(in, os.Stdout); !reflect.DeepEqual(c, test.counts) || !reflect.DeepEqual(u, test.utflen) {
			t.Errorf("Charcount(%q) = %v , %v", test.input, c, u)
		}
	}
}
