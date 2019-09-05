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

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(10)
	x.Add(9)
	if x.Len() != 4 {
		t.Error()
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(9)
	x.Add(144)
	x.Clear()
	if x.String() != "{}" {
		t.Error()
	}
	if x.Len() != 0 {
		t.Error()
	}
}
func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(9)
	x.Add(144)
	x.Remove(9)
	if x.String() != "{1 144}" {
		t.Error()
	}
	x.Remove(1)
	if x.String() != "{144}" {
		t.Error()
	}
	x.Remove(144)
	if x.String() != "{}" {
		t.Error()
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(9)
	x.Add(144)

	y := x.Copy()
	x.Add(145)
	if y.String() != "{1 9 144}" {
		t.Error()
	}

}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet

	x.AddAll(1, 9, 144)
	y.AddAll(2, 9, 100)
	x.IntersectWith(&y)
	if x.String() != "{9}" {
		t.Error()
	}

	x.Clear()
	y.Clear()
	x.AddAll(144)
	y.AddAll(10)
	x.IntersectWith(&y)
	if x.String() != "{}" {
		t.Error()
	}

	x.Clear()
	y.Clear()
	x.AddAll(1, 2, 3, 144)
	x.IntersectWith(&y)
	if x.String() != "{}" {
		t.Error()
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet

	x.AddAll(1, 9, 144)
	y.AddAll(2, 9, 100)
	x.DifferenceWith(&y)
	if x.String() != "{1 144}" {
		t.Error()
	}

	x.Clear()
	y.Clear()
	x.AddAll(1, 2, 9)
	y.AddAll(2, 9, 255)
	x.DifferenceWith(&y)
	if x.String() != "{1}" {
		t.Error()
	}

	x.Clear()
	y.Clear()
	x.AddAll(1, 9, 255)
	y.AddAll(1, 9, 255, 4)
	x.DifferenceWith(&y)
	if x.String() != "{}" {
		t.Error()
	}

}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet

	x.AddAll(1, 9, 144)
	y.AddAll(2, 9, 100)
	x.SymmetricDifference(&y)
	if x.String() != "{1 2 100 144}" {
		t.Error()
	}

	x.Clear()
	y.Clear()
	x.AddAll(1, 2, 9)
	y.AddAll(2, 9, 255)
	x.SymmetricDifference(&y)
	if x.String() != "{1 255}" {
		t.Error()
	}

	x.Clear()
	y.Clear()
	x.AddAll(1, 9, 255)
	y.AddAll(1, 9, 255, 4)
	x.SymmetricDifference(&y)
	if x.String() != "{4}" {
		t.Error()
	}

	x.Clear()
	y.Clear()
	x.AddAll(1, 9, 255, 4)
	y.AddAll(1, 9, 255, 4)
	x.SymmetricDifference(&y)
	if x.String() != "{}" {
		t.Error()
	}

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
