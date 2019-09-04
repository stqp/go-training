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

type Cases struct {
	arg      []string
	expected int
}

func TestElementByTagName(t *testing.T) {

	cases := []Cases{
		{arg: []string{"html"}, expected: 1},
		{arg: []string{"head"}, expected: 1},
		{arg: []string{"div"}, expected: 7},
		{arg: []string{"html", "a", "span"}, expected: 10},
	}

	r := strings.NewReader(body)
	doc, _ := html.Parse(r)

	for _, c := range cases {
		actual := ElementsByTagName(doc, c.arg...)
		if len(actual) != c.expected {
			t.Error("expected:", c.expected, "actual:", actual)
		}
	}
}
