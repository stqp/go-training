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
	<link rel="stylesheet" href="http://ajax.googleapis.com/ajax/libs/jqueryui/1.8.19/themes/redmond/jquery-ui.css">
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
			<img src="https://gooogle.com/image1.png">
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
	links := visit(nil, doc)
	ans := []string{
		"https://google.com/cdn/jqeury.js",
		"http://ajax.googleapis.com/ajax/libs/jqueryui/1.8.19/themes/redmond/jquery-ui.css",
		"https://google.com/?q=trending",
		"https://google.com/?q=trending2",
		"https://google.com/?q=trending3",
		"https://gooogle.com/image1.png",
		"https://www.test.com/about",
		"https://www.test.com/products",
		"https://www.test.com/history",
	}
	for i, link := range links {
		if ans[i] != link {
			t.Error("Expected:", i, ans[i], "Actual:", link)
		}
	}

}
