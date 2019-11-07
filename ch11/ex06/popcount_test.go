package popcount

import (
	"testing"
)

func benchmarkPopCount(b *testing.B, d uint64) {
	for i := 0; i < b.N; i++ {
		PopCount(0xFFFFF)
	}
}
func benchmarkPopCountByClearing(b *testing.B, d uint64) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(d)
	}
}
func benchmarkPopCountByShifting(b *testing.B, d uint64) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(d)
	}
}

func BenchmarkPopCount_0xF1(b *testing.B)           { benchmarkPopCount(b, 0xF) }
func BenchmarkPopCountByClearing_0xF1(b *testing.B) { benchmarkPopCountByClearing(b, 0xF) }
func BenchmarkPopCountByShifting_0xF1(b *testing.B) { benchmarkPopCountByShifting(b, 0xF) }

func BenchmarkPopCount_0xF5(b *testing.B)           { benchmarkPopCount(b, 0xFFFF) }
func BenchmarkPopCountByClearing_0xF5(b *testing.B) { benchmarkPopCountByClearing(b, 0xFFFF) }
func BenchmarkPopCountByShifting_0xF5(b *testing.B) { benchmarkPopCountByShifting(b, 0xFFFF) }

func BenchmarkPopCount_0xF10(b *testing.B)           { benchmarkPopCount(b, 0xFFFFFFFFFF) }
func BenchmarkPopCountByClearing_0xF10(b *testing.B) { benchmarkPopCountByClearing(b, 0xFFFFFFFFFF) }
func BenchmarkPopCountByShifting_0xF10(b *testing.B) { benchmarkPopCountByShifting(b, 0xFFFFFFFFFF) }

func BenchmarkPopCount_0xF16(b *testing.B) {
	benchmarkPopCount(b, 0xFFFFFFFFFFFFFFFF)
}
func BenchmarkPopCountByClearing_0xF16(b *testing.B) {
	benchmarkPopCountByClearing(b, 0xFFFFFFFFFFFFFFFF)
}
func BenchmarkPopCountByShifting_0xF16(b *testing.B) {
	benchmarkPopCountByShifting(b, 0xFFFFFFFFFFFFFFFF)
}
