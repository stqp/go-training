package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountv4(0xFFFFFFFFFFFFFFFF)
	}
}

func TestPopCountv4(t *testing.T) {
	res := PopCountv4(0xFFFFFFFFFFFFFFFF)
	if res != 64 {
		t.Errorf("res is %d", res)
	}
	res = PopCountv4(0xF)
	if res != 4 {
		t.Errorf("res is %d", res)
	}
}
