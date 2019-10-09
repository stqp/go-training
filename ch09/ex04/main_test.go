package main

import (
	"testing"
)

func BenchmarkPipelineStage100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := Pipe{}
		p.build(100000)
	}
}

func BenchmarkPipelineStage1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := Pipe{}
		p.build(1000000)
	}
}

func SkipBecauseTakesTooLongTime_BenchmarkPipelineStage5000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := Pipe{}
		p.build(5000000)
	}
}

func BenchmarkPassThrough1000(b *testing.B) {
	b.StopTimer()
	p := Pipe{}
	p.build(1000)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.passThrough(99)
	}
}

func TestPassThrough(t *testing.T) {
	p := Pipe{}
	p.build(100)

	// まだパイプには何も流れていないので0のはず
	if p.passedCount != 0 {
		t.Error()
	}

	if p.passThrough(99) != 99 {
		t.Error()
	}

	// 通ったパイプの総数をチェック
	if p.passedCount != 100 {
		t.Error()
	}
}

func BenchmarkPassThrough100000(b *testing.B) {
	b.StopTimer()
	p := Pipe{}
	p.build(100000)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		p.passThrough(99)
	}
}
