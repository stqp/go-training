#! /bin/bash

go test

# [01011789@CA1525 ex05]$ go test         
# PASS
# ok  	github.com/stqp/go-training/ch08/ex05	4.785s
# [01011789@CA1525 ex05]$

go test -bench=.

# [01011789@CA1525 ex05]$ go test -bench=.
# goos: darwin
# goarch: amd64
# pkg: github.com/stqp/go-training/ch08/ex05
# BenchmarkHeavy-4           	       1	26778931256 ns/op
# BenchmarkConcurrency1-4    	       1	26892631152 ns/op
# BenchmarkConcurrency2-4    	       1	16449188337 ns/op
# BenchmarkConcurrency4-4    	       1	15584230170 ns/op
# BenchmarkConcurrency8-4    	       1	15708460589 ns/op
# BenchmarkConcurrency16-4   	       1	15803974606 ns/op
# BenchmarkConcurrency64-4   	       1	15374958752 ns/op
# PASS
# ok  	github.com/stqp/go-training/ch08/ex05	175.497s
# 
# [01011789@CA1525 ~]$ sysctl hw.physicalcpu hw.logicalcpu
# hw.physicalcpu: 2
# hw.logicalcpu: 4
