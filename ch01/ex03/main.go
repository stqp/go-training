package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func buildData(size int) []string {
	bigArgs := []string{}
	for i := 0; i < size; i++ {
		bigArgs = append(bigArgs, "aiueo")
	}
	return bigArgs
}

type fn func([]string)

func oldOne(data []string) {
	var s, sep string
	for i := 1; i < len(data); i++ {
		s += sep + data[i]
		sep = " "
	}
}

func newOne(data []string) {
	strings.Join(data, " ")
}

func letsKeisoku(title string, f fn, data []string) {
	start := time.Now()
	f(data)
	secs := time.Since(start).Seconds()
	fmt.Printf("  %s : %.2fs, data size\n", title, secs)
}

func main() {
	for i := 10; i < 18; i++ {
		data := buildData(int(math.Pow(2.0, float64(i))))
		fmt.Printf("When data size is : %d\n", len(data))
		letsKeisoku("oldOne", oldOne, data)
		letsKeisoku("newOne", newOne, data)
		fmt.Println("")
	}
}
