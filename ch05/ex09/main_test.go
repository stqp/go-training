package main

import (
	"testing"
)

var (
	body = `$foo is a $foo, but $foo`
)

func TestMain(t *testing.T) {
	actual := expand(body, func(s string) string { return "cat" })
	expected := "cat is a cat, but cat"
	if actual != expected {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}
