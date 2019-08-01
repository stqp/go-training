package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordreq(counts map[string]map[rune]int, category string, r rune) {
	if counts[category] == nil {
		counts[category] = make(map[rune]int)
	}
	counts[category][r]++
}

func main() {
	file, err := os.Open("testdata")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	freq := make(map[string]int)
	for scanner.Scan() {
		freq[strings.ToLower(scanner.Text())]++
	}
	for k, v := range freq {
		fmt.Printf("'%s' , %d \n", k, v)
	}

}
