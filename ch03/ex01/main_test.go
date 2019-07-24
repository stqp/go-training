package main

import (
	"math"
	"testing"
)

func TestIsFinite(t *testing.T) {
	data := []float64{1, 2, 3}
	if isFinite(data) == false {
		t.Error("Errror", data)
	}
	data = []float64{1, math.Inf(0)}
	if isFinite(data) == true {
		t.Error("Errror", data)
	}
}
