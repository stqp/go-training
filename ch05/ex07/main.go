package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func main() {}

func outline(doc *html.Node) error {
	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attrs := ""
		for _, a := range n.Attr {
			attrs += " " + a.Key + "='" + a.Val + "'"
		}

		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrs)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrs)
			depth++
		}

	} else if n.Type == html.TextNode && n.Data[0] != 10 {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			return
		} else {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}

	}
}
