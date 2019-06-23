package main

import (
	"fmt"
	"os"
	"strconv"

	tempconv "github.com/stqp/go-training/ch02/ex01"
)

func usage() {
	fmt.Println("please specify 2 arguments")
	fmt.Println("for example:")
	fmt.Println(" $ go run main.go [temp | length | weight] <data>")
}

func main() {
	if len(os.Args) == 2 {
		usage()
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
