package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("# Variant must be uppercase letter (A ~ Z) #\n ")
}

func main() {
	usage()

	fmt.Print("Expr: ")
	s := bufio.NewScanner(os.Stdin)
	ok := s.Scan()
	if !ok {
		panic(fmt.Sprintf("%s", s.Err()))
	}
	in := s.Text()
	expr, err := Parse(in)
	if err != nil {
		fmt.Printf("expr is invalid: %s\n", err)
		return
	}
	env := Env{}
	for _, v := range in {
		if 'A' <= v && v <= 'Z' {
			fmt.Printf("%s: ", string(v))
			var num string
			fmt.Scan(&num)
			f, err := strconv.ParseFloat(num, 64)
			if err != nil {
				fmt.Printf("variant value is invalid: %s\n", num)
				return
			}
			env[Var(string(v))] = f
		}
	}
	fmt.Println("= ", expr.Eval(env))
}
