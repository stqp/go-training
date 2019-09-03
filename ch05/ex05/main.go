package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var (
	t map[string]int
)

func init() {
	t = make(map[string]int)
}
func main() {}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return 0, 0
	}
	ws, is := 0, 0
	if n.Type == html.TextNode {
		ws = len(strings.Fields(n.Data))
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		is = 1
	}
	a, b := countWordsAndImages(n.FirstChild)
	c, d := countWordsAndImages(n.NextSibling)
	ws += a + c
	is += b + d
	return ws, is
}
