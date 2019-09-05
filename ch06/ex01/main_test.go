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
	fmt.Println()
}
func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	if x.String() != "{1}" {
		t.Error()
	}
	x.Add(144)
	x.Add(9)
	if x.String() != "{1 9 144}" {
		t.Error()
	}
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
