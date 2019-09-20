package main

import (
	"io"

	"golang.org/x/net/html"
)

func main() {}

// Reader is reader
type Reader struct {
	s string
	i int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(p, r.s[r.i:])
	r.i += int64(n)
	return
}

// NewReader can read string
func NewReader(s string) *Reader {
	return &Reader{s, 0}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
