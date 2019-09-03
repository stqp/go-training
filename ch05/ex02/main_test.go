package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var (
	body = `
	<html>

<head>
    <script type="text/javascript" src="https://google.com/cdn/jqeury.js"></script>
</head>

<body>
    <div class="content">
        <p>this is content</p>
    </div>
    <div id="body">
        <h1>News</h1>
        <h2>Pick Ups</h2>
        <div id="pickups_news">
            <div class="news">
                <span>news1</span>
                <a href="https://google.com/?q=trending">new1_link</a>
            </div>
            <div class="news">
                <span>news2</span>
                <a href="https://google.com/?q=trending2">new2_link</a>
            </div>
            <div class="news">
                <span>news3</span>
                <a href="https://google.com/?q=trending3">new3_link</a>
            </div>
        </div>
    </div>
    <div class="footer">
        <a href="https://www.test.com/about">about</a>
        <a href="https://www.test.com/products">products</a>
        <a href="https://www.test.com/history">history</a>
    </div>
</body>

</html>
`
)

func TestMain(t *testing.T) {
	r := strings.NewReader(body)
	doc, _ := html.Parse(r)
	tags := make(table, 1)
	tags = tagCount(tags, doc)
	ans := table{
		"html":   1,
		"body":   1,
		"p":      1,
		"head":   1,
		"script": 1,
		"div":    7,
		"h1":     1,
		"h2":     1,
		"span":   3,
		"a":      6,
	}
	for i, c := range tags {
		if ans[i] != c {
			t.Error("Expected:", i, ans[i], "Actual:", c)
		}
	}

}
