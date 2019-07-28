package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		t = flag.String("t", "256", "string flag")
	)
	flag.Parse()

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		in := stdin.Text()
		switch *t {
		case "256":
			fmt.Printf("%x\n", sha256.Sum256([]byte(in)))
		case "384":
			fmt.Printf("%x\n", sha512.Sum384([]byte(in)))
		case "512":
			fmt.Printf("%x\n", sha512.Sum512([]byte(in)))
		}
	}
}
