package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func clean() {
	dir := outbase()
	if err := os.RemoveAll(dir); err != nil {
		fmt.Println(err)
	}
}

func outbase() string {
	return "./out"
}

func exists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func outDir(urlstr string) string {
	base := outbase()
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}

	out := base + "/" + u.Hostname()
	if !exists(out) {
		if err := os.MkdirAll(out, 0777); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return out
}

func fetch(url string, out string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	local = out + "/" + local
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

func copyWebPage(url string) {
	out := outDir(url)
	f, n, err := fetch(url, out)
	if err != nil {
		fmt.Println(f, n, err)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	copyWebPage(url)
	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	clean()
	breadthFirst(crawl, os.Args[1:])
}

func isSubDomain(subURL, mainURL string) bool {
	sub, err := url.Parse(subURL)
	if err != nil {
		log.Fatal(err)
	}
	main, err := url.Parse(mainURL)
	if err != nil {
		log.Fatal(err)
	}
	subhost := sub.Hostname()
	mainhost := main.Hostname()

	maindomain := mainhost[strings.Index(mainhost, ".")+1:]
	return strings.Contains(subhost, maindomain)
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if isSubDomain(link.String(), url) {
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
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
