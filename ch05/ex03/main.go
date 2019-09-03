package main

import (
	"fmt"

	"golang.org/x/net/html"
)

type table map[string]int

func main() {}

func printTagBody(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
		fmt.Println(n.Data, ":")
	}
	if n.Type == html.TextNode {
		if n.Data != "" {
			fmt.Println(n.Data)
			fmt.Println("----------------")
		}
	}
	printTagBody(n.FirstChild)
	printTagBody(n.NextSibling)

}
