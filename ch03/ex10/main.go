package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func reverse(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return b
}

func comma(s string) string {
	var b bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		if i > 0 && i%3 == 2 && i < n-1 {
			b.Write([]byte{s[n-1-i], ','})
		} else {
			b.Write([]byte{s[n-1-i]})
		}
	}
	return string(reverse(b.Bytes()))
}
