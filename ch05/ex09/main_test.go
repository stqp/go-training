package main

import (
	"testing"
)

var (
	body = `This is a $foo.`
)

func TestMain(t *testing.T) {
	actual := expand(body, func(s string) string { return "****" })
	expected := "This is a ****."
	if actual != expected {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}
