package main

import (
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	actual := expand("hello  $world now.", strings.ToUpper)
	expected := "WORLD"
	if actual != expected {
		t.Error("Expected:", expected, "Actual:", actual)
	}
}
