package main

import (
	"time"
)

func main() {}

func oneSecondTalk() int64 {
	var count int64
	in1 := make(chan string)
	in2 := make(chan string)
	done := make(chan struct{})
	go func() {
		for msg := range in1 {
			count++
			in2 <- msg
		}
	}()
	go func() {
		for msg := range in2 {
			count++
			in1 <- msg
		}
	}()
	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	in1 <- "let's start"
	<-done
	return count
}
