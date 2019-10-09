package main

import (
	"bytes"
	"testing"
)

type nullWriter int

func (w nullWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func Test(t *testing.T) {
	var s1 string
	w := bytes.NewBufferString(s1)
	runHeavyMandelbrot(w)

	var s2 string
	w2 := bytes.NewBufferString(s1)
	runConcurrencyMandelbrot(w2, 1000)

	if s1 != s2 {
		t.Error()
	}
}

func BenchmarkHeavy(b *testing.B) {
	w := nullWriter(0)
	for n := 0; n < b.N; n++ {
		runHeavyMandelbrot(w)
	}
}

func BenchmarkConcurrency1(b *testing.B) {
	w := nullWriter(0)
	for n := 0; n < b.N; n++ {
		runConcurrencyMandelbrot(w, 1)
	}
}

func BenchmarkConcurrency2(b *testing.B) {
	w := nullWriter(0)
	for n := 0; n < b.N; n++ {
		runConcurrencyMandelbrot(w, 2)
	}
}

func BenchmarkConcurrency4(b *testing.B) {
	w := nullWriter(0)
	for n := 0; n < b.N; n++ {
		runConcurrencyMandelbrot(w, 4)
	}
}

func BenchmarkConcurrency8(b *testing.B) {
	w := nullWriter(0)
	for n := 0; n < b.N; n++ {
		runConcurrencyMandelbrot(w, 8)
	}
}

func BenchmarkConcurrency16(b *testing.B) {
	w := nullWriter(0)
	for n := 0; n < b.N; n++ {
		runConcurrencyMandelbrot(w, 16)
	}
}

func BenchmarkConcurrency64(b *testing.B) {
	w := nullWriter(0)
	for n := 0; n < b.N; n++ {
		runConcurrencyMandelbrot(w, 64)
	}
}
