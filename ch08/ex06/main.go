package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type Work struct {
	link  string
	depth int
}

var maxdepth int

func main() {
	flag.IntVar(&maxdepth, "depth", 0, "max depth")
	flag.Parse()

	worklist := make(chan []Work)
	unseenLinks := make(chan Work)

	go func() {
		var works []Work
		for _, link := range os.Args[2:] {
			works = append(works, Work{link: link, depth: 0})
		}
		worklist <- works
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for w := range unseenLinks {
				foundLinks := crawl(w.link)
				var works []Work
				for _, link := range foundLinks {
					works = append(works, Work{link: link, depth: w.depth + 1})
				}
				go func() { worklist <- works }()
			}
		}()
	}

	seen := make(map[string]bool)
	for works := range worklist {
		for _, w := range works {
			if !seen[w.link] && w.depth < maxdepth {
				seen[w.link] = true
				unseenLinks <- w
			}
		}
	}
}
