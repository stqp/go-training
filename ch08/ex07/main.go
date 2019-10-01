package main

import (
	"flag"
	"fmt"
)

var out string
var hostname = "localhost"
var port = "5000"

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "crawl", "crawl is 1 , server is 2")
	flag.StringVar(&out, "out", "./out", "output directory")
	flag.Parse()

	if mode == "1" {
		startCrawl(flag.Args())
	} else if mode == "2" {
		serve()
	} else {
		fmt.Printf("invalid mode number : %s\n", mode)
	}

}
