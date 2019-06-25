package main

import (
	"fmt"

	"github.com/stqp/go-training/ch02/ex01/tempconv"
)

func main() {

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	var k tempconv.Kelvin

	c = 100
	f = tempconv.CToF(c)
	k = tempconv.CToK(c)

	fmt.Println(c.String())
	fmt.Println(f.String())
	fmt.Println(k.String())

}
