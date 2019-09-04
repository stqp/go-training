package main

import "testing"

func Test(t *testing.T) {
	if a() != 10 {
		t.Error("invalid")
	}
}
