package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFFFFFFFFFFFFF)
	}
}

func BenchmarkPopCountv3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountv3(0xFFFFFFFFFFFFFFFF)
	}
}

func TestPopCountv3(t *testing.T) {
	res := PopCountv3(0xFFFFFFFFFFFFFFFF)
	if res != 64 {
		t.Errorf("res is %d", res)
	}
	res = PopCountv3(0xF)
	if res != 4 {
		t.Errorf("res is %d", res)
	}
}
