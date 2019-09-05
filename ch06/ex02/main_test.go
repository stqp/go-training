// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3, 4, 5)
	if x.String() != "{1 2 3 4 5}" {
		t.Error()
	}
}
