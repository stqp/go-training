package main

import "testing"

func Test(t *testing.T) {
	if a() != 1 {
		t.Error("invalid")
	}
}
