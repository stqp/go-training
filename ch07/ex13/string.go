package eval

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
