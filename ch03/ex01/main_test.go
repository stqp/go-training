package main

import (
	"testing"
)

func TestIsFinite(t *testing.T) {
	data := []float64{1, 2, 3}
	if isFinite(data) != false {
		t.Error("Errror", data)
	}
}
