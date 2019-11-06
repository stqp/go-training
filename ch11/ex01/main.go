package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func Charcount(in io.Reader, out io.Writer) {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	r := bufio.NewReader(in)
	for {
		r, n, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Fprintf(out, "rune\tcount\n")
	for c, n := range counts {
		fmt.Fprintf(out, "%q\t%d\n", c, n)
	}
	fmt.Fprint(out, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(out, "%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Fprintf(out, "\n%d invalid UTF-8 characters\n", invalid)
	}
}

func main() {
	Charcount(os.Stdin, os.Stdout)
}
