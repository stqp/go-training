package main

import (
	"strings"
)

func main() {}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	for _, w := range words {
		i := strings.Index(w, "$")
		if i >= 0 {
			return f(w[i+1:])
		}
	}
	return ""
}
