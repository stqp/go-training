package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	xmlselect(os.Stdin)
}

func xmlselect(in io.Reader) {
	dec := xml.NewDecoder(in)
	var stack []Element
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
			elem := Element{tok.Name, tok.Attr, nil}
			stack = append(stack, elem)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			//stack = append(stack, CharData(tok))
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
