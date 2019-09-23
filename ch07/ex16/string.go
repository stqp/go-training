package main

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	a := string(u.op) + u.x.String()
	return a
}

func (b binary) String() string {
	a := "(" + b.x.String() + string(b.op) + b.y.String() + ")"
	return a
}

func (c call) String() string {
	a := c.fn + "("
	if c.fn == "pow" {
		a += c.args[0].String() + "," + c.args[1].String()
	} else {
		a += c.args[0].String()
	}
	a += ")"
	return a
}

func (l list) String() string {
	a := l.fn + "["
	if len(l.args) != 0 {
		a += l.args[0].String()
		for _, v := range l.args[1:] {
			a += "," + v.String()
		}
	}
	a += "]"
	return a
}
