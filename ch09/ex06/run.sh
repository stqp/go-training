#! /bin/bash

# まず実行環境の確認。いまのMacは物理2コアしかない↓
# 
# [01011789@CA1525 ~]$ sysctl hw.physicalcpu hw.logicalcpu
# hw.physicalcpu: 2
# hw.logicalcpu: 4


GOMAXPROCS=1 go test -bench=.
# [01011789@CA1525 ex06]$ GOMAXPROCS=1 go test -bench=.
# goos: darwin
# goarch: amd64
# pkg: github.com/stqp/go-training/ch09/ex06
# BenchmarkHeavy         	       1	2749386718 ns/op
# BenchmarkConcurrency1  	       1	2788866667 ns/op
# BenchmarkConcurrency2  	       1	2763257208 ns/op
# BenchmarkConcurrency4  	       1	2765308800 ns/op
# BenchmarkConcurrency8  	       1	2783415777 ns/op
# BenchmarkConcurrency16 	       1	2981408221 ns/op
# BenchmarkConcurrency64 	       1	2967622086 ns/op
# PASS
# ok  	github.com/stqp/go-training/ch09/ex06	25.638s
# [01011789@CA1525 ex06]$
# ↑　ほとんど時間に差がない。これは1コアしか使えず並列化できていないってことかな。


GOMAXPROCS=2 go test -bench=.
# [01011789@CA1525 ex06]$ GOMAXPROCS=2 go test -bench=.
# goos: darwin
# goarch: amd64
# pkg: github.com/stqp/go-training/ch09/ex06
# BenchmarkHeavy-2           	       1	2777655904 ns/op
# BenchmarkConcurrency1-2    	       1	2844950823 ns/op
# BenchmarkConcurrency2-2    	       1	1941312467 ns/op
# BenchmarkConcurrency4-2    	       1	1742550719 ns/op
# BenchmarkConcurrency8-2    	       1	1678589691 ns/op
# BenchmarkConcurrency16-2   	       1	1717917124 ns/op
# BenchmarkConcurrency64-2   	       1	1755513907 ns/op
# PASS
# ok  	github.com/stqp/go-training/ch09/ex06	19.070s
# [01011789@CA1525 ex06]$
#
# ↑ をみると「BenchmarkConcurrency2」以降高速になっている。
# だがコア上限に引っかかり並列化の設定値を増やしても早くならないってことかな。