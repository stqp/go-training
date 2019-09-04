package main

import (
	"golang.org/x/net/html"
)

func main() {}

var t []*html.Node

// ElementsByTagName is ElementsByTagName.
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	t = nil
	return eelementsByTagName(doc, name...)
}

func eelementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if doc == nil {
		return nil
	}
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				t = append(t, doc)
			}
		}
	}
	eelementsByTagName(doc.FirstChild, name...)
	eelementsByTagName(doc.NextSibling, name...)
	return t
}
