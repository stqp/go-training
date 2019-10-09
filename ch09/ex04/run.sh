#! /bin/bash
go test
go test -bench=.

#[01011789@CA1525 ex04]$ go test -bench=.
#goos: darwin
#goarch: amd64
#pkg: github.com/stqp/go-training/ch09/ex04
#BenchmarkPipelineStage100000-4    	       3	 368426776 ns/op
#BenchmarkPipelineStage1000000-4   	       1	3236174828 ns/op
#BenchmarkPassThrough1000-4        	    5000	    351267 ns/op
#BenchmarkPassThrough100000-4      	      30	  45113729 ns/op
#PASS
#ok  	github.com/stqp/go-training/ch09/ex04	19.939s