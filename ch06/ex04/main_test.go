// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func debug(x IntSet) {
	for _, v := range x.words {
		fmt.Printf("[%b] ", v)
	}
	fmt.Printf(x.String())
	fmt.Println()
}

func TestElems(t *testing.T) {
	var x IntSet
	x.AddAll(1, 9, 144)

	if len(x.Elems()) != 3 {
		t.Error()
	}

	expected := []int{1, 9, 144}
	for i, e := range x.Elems() {
		if e != expected[i] {
			t.Error()
		}
	}

}
