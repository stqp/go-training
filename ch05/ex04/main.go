package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var (
	t map[string]string
)

func init() {
	t = make(map[string]string)
	t["a"] = "href"
	t["img"] = "src"
	t["script"] = "src"
	t["link"] = "href"
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		key := t[n.Data]
		if key != "" {
			for _, a := range n.Attr {
				if a.Key == key {
					links = append(links, a.Val)
				}
			}
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
