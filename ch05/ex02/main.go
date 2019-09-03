package main

import (
	"golang.org/x/net/html"
)

type table map[string]int

func main() {}

func tagCount(t table, n *html.Node) table {
	if n == nil {
		return t
	}
	if n.Type == html.ElementNode {
		t[string(n.Data)]++
	}
	t = tagCount(t, n.FirstChild)
	t = tagCount(t, n.NextSibling)
	return t
}
