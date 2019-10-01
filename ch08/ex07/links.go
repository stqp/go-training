package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	u "net/url"
	"strings"

	"golang.org/x/net/html"
)

func extract(url string) ([]string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("fail to read http response body"))
	}

	r := strings.NewReader(string(content))
	doc, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	r2 := strings.NewReader(string(content))
	doc2, err := html.Parse(r2)
	u2, _ := u.Parse(url)

	mirror2(u2.Path, doc2)

	var links []string
	visitNode := func(n *html.Node, w io.Writer) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}

				if err != nil {
					panic(fmt.Sprint("fail to get hostname"))
				}

				if isSameDomain(url, a.Val) {
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, w io.Writer), w io.Writer) {
	if pre != nil {
		pre(n, w)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, w)
	}
	if post != nil {
		post(n, w)
	}
}

func isSameDomain(url1 string, url2 string) bool {
	u1, err := url.Parse(url1)
	if err != nil {
		panic(fmt.Sprintf("can't parse url: %s", url1))
	}
	u2, err := url.Parse(url2)
	if err != nil {
		panic(fmt.Sprintf("can't parse url: %s", url2))
	}
	return u1.Hostname() == "" || u2.Hostname() == "" || u1.Hostname() == u2.Hostname()
}
