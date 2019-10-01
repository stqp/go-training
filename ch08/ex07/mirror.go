package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

//func mirror(r *http.Response) {
func mirror(path string, content []byte) {
	if strings.HasSuffix(path, "/") || path == "" {
		path += "index.html"
	}

	if !strings.Contains(filepath.Base(path), ".") {
		path += ".html"
	}
	path = out + "/" + path
	path, _ = filepath.Abs(path)
	ifnotExistsThenCreateDir(path)
	writeToFile(path, content)
}

func visitWithReplace(doc *html.Node, hostname string) string {
	t := make(map[string]string)
	t["a"] = "href"
	t["img"] = "src"
	t["script"] = "src"
	t["link"] = "href"

	var res string
	var visit func(n *html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode {

			key := t[n.Data]
			if key != "" {
				for i, a := range n.Attr {
					if a.Key == key {
						//fmt.Println("a.Key:", a.Key, "a.Val:", a.Val, "key:", key, "n.Attr:", n.Attr)
						newKey, err := url.Parse(a.Val)
						if err != nil {
							// ignore error.
							// for example:
							//: net/url: invalid control character in URL
							continue
						}
						if newKey.Host != "" {
							newKey.Host = hostname
						}
						n.Attr[i].Val = newKey.String()
					}
				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)
	return res
}

func mirror2(path string, doc *html.Node) {
	if strings.HasSuffix(path, "/") || path == "" {
		path += "index.html"
	}

	if !strings.Contains(filepath.Base(path), ".") {
		path += ".html"
	}
	path = out + "/" + path
	path, _ = filepath.Abs(path)
	ifnotExistsThenCreateDir(path)

	visitWithReplace(doc, hostname)

	var w strings.Builder
	nodeToString(doc, &w)
	writeToFile(path, []byte(w.String()))
}

func nodeToString(doc *html.Node, w io.Writer) {
	forEachNode(doc, startElement, endElement, w)
}

var cnt string

//var depth int

func startElement(n *html.Node, w io.Writer) {
	if n.Type == html.ElementNode {
		attrs := ""
		for _, a := range n.Attr {
			attrs += " " + a.Key + "='" + a.Val + "'"
		}

		if n.FirstChild == nil {
			w.Write([]byte(fmt.Sprintf("<%s%s></%s>\n", n.Data, attrs, n.Data)))
		} else {
			w.Write([]byte(fmt.Sprintf("<%s%s>", n.Data, attrs)))
		}

	} else if n.Type == html.TextNode {
		w.Write([]byte(fmt.Sprintf("%s", n.Data)))
	} else if n.Type == html.CommentNode {
		w.Write([]byte(fmt.Sprintf("<!-- %s -->", n.Data)))
	}
}

func endElement(n *html.Node, w io.Writer) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			return
		} else {
			//depth--
			//cnt += fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
			w.Write([]byte(fmt.Sprintf("</%s>\n", n.Data)))
		}
	}
}

func writeToFile(path string, content []byte) {
	f, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintf("failed to create output file : %s", path))
	}
	defer f.Close()
	f.Write(content)
}
