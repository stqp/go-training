package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CountsInfo have information related to counts.
type CountsInfo struct {
	counts  int
	foundAt []string
}

// Counts is counts.
type Counts map[string]*CountsInfo

func main() {
	counts := make(Counts)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for key, info := range counts {
		if info.counts > 1 {
			foundAt := strings.Join(info.foundAt, " ")
			fmt.Printf("%d\t%s\t%s\n", info.counts, key, foundAt)
		}
	}
}

func countLines(f *os.File, counts Counts) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = &CountsInfo{}
		}

		counts[input.Text()].counts++
		already := false
		for _, v := range counts[input.Text()].foundAt {
			if v == f.Name() {
				already = true
			}
		}
		if !already {
			counts[input.Text()].foundAt = append(counts[input.Text()].foundAt, f.Name())
		}

	}
}
