package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	xmlselect(os.Stdin)
}

func xmlselect(in io.Reader) {
	dec := xml.NewDecoder(in)
	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", stringsJoin(stack, " "), tok)
			}
		}
	}
}

func stringsJoin(stack []xml.StartElement, sep string) string {
	ans := ""
	for _, elem := range stack {
		ans += elem.Name.Local
		if len(elem.Attr) > 0 {
			ans += "["
			ans += elem.Attr[0].Name.Local + "="
			ans += elem.Attr[0].Value
			for _, v := range elem.Attr[1:] {
				ans += "," + v.Name.Local + "="
				ans += v.Value
			}
			ans += "]"
		}
		ans += sep
	}
	return ans
}

func containsAll(stack []xml.StartElement, selectors []string) bool {
	for len(selectors) <= len(stack) {
		if len(selectors) == 0 {
			return true
		}
		// ex: div[id=test,class=wide]
		if i := strings.Index(selectors[0], "["); i > 0 {
			tag := selectors[0][:i]
			attrsStr := selectors[0][i+1 : len(selectors[0])-1]
			attrs := strings.Split(attrsStr, ",")

			matched := true
			if tag != stack[0].Name.Local {
				matched = false
			}
			for _, attr := range attrs {
				k := strings.Split(attr, "=")[0]
				v := strings.Split(attr, "=")[1]
				found := false
				for _, v2 := range stack[0].Attr {
					if v2.Name.Local == k && v2.Value == v {
						found = true
					}
				}
				if !found {
					matched = false
					break
				}
			}
			if matched {
				selectors = selectors[1:]
			}

		} else {
			if stack[0].Name.Local == selectors[0] {
				selectors = selectors[1:]
			}
		}
		stack = stack[1:]
	}
	return false
}
