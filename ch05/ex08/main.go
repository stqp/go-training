package main

import (
	"golang.org/x/net/html"
)

func main() {}

// ElementByID is ElementByID.
func ElementByID(n *html.Node, id string) *html.Node {
	return forEachNode(n, element, id)
}

func forEachNode(n *html.Node, f func(n *html.Node, id string) bool, id string) *html.Node {
	if n == nil {
		return nil
	}
	if f != nil && f(n, id) == true {
		return n
	}
	a := forEachNode(n.FirstChild, f, id)
	if a != nil {
		return a
	}
	b := forEachNode(n.NextSibling, f, id)
	if b != nil {
		return b
	}
	return nil
}

func element(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}
