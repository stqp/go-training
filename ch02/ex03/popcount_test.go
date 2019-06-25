package popcount

import (
	"testing"
	//"github.com/stqp/go-training/ch02/ex03/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountv2(0xFFFFFFFFFFFFFFFF)
	}
}
