package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unsafe"
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
	var sign string
	if s[0] == '-' {
		sign = s[:1]
		s = s[1:]
	}

	dotIndex := strings.Index(s, ".")
	var low string
	if dotIndex > 0 {
		low = s[dotIndex:]
		s = s[:dotIndex]
	}
	var b bytes.Buffer
	for i := len(low) - 1; i >= 0; i-- {
		b.Write([]byte{low[i]})
	}
	n := len(s)
	for i := 0; i < n; i++ {
		if i > 0 && i%3 == 2 && i < n-1 {
			b.Write([]byte{s[n-1-i], ','})
		} else {
			b.Write([]byte{s[n-1-i]})
		}
	}
	b.Write(*(*[]byte)(unsafe.Pointer(&sign)))
	return string(reverse(b.Bytes()))
}
