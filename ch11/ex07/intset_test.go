package intset

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkIntSetAdd(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	set := IntSet{}
	for i := 0; i < b.N; i++ {
		n := rng.Intn(math.MaxInt32)
		set.Add(n)
	}
}

func BenchmarkMapSetAdd(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	set := MapSet{}
	for i := 0; i < b.N; i++ {
		n := rng.Intn(math.MaxInt32)
		set.Add(n)
	}
}

func BenchmarkIntSetUnionWith(b *testing.B) {

	b.StopTimer()

	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	set := IntSet{}
	set2 := IntSet{}
	for i := 0; i < b.N; i++ {
		n := rng.Intn(math.MaxInt32)
		set.Add(n)
	}
	for i := 0; i < b.N; i++ {
		n := rng.Intn(math.MaxInt32)
		set2.Add(n)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.UnionWith(&set2)
	}
}

func BenchmarkMapSetUnionWith(b *testing.B) {

	b.StopTimer()

	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	set := MapSet{}
	set2 := MapSet{}
	for i := 0; i < b.N; i++ {
		n := rng.Intn(math.MaxInt32)
		set.Add(n)
	}
	for i := 0; i < b.N; i++ {
		n := rng.Intn(math.MaxInt32)
		set2.Add(n)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		set.UnionWith(&set2)
	}
}

// $ go test -bench=.
// goos: darwin
// goarch: amd64
// pkg: github.com/stqp/go-training/ch11/ex07
// BenchmarkIntSetAdd-4         	       5	 204147884 ns/op
// BenchmarkMapSetAdd-4         	 5000000	       302 ns/op
// BenchmarkIntSetUnionWith-4   	       2	 546954559 ns/op
// BenchmarkMapSetUnionWith-4   	   10000	    706250 ns/op
// PASS
// ok  	github.com/stqp/go-training/ch11/ex07	15.817s
// $
