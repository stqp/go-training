package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Package struct {
	ImportPath string
	Imports    []string
	Name       string
	Deps       []string
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("specify package name.")
		os.Exit(1)
	}
	target := os.Args[1]

	args := []string{"go", "list", "-e", "..."}
	out, err := exec.Command(args[0], args[1:]...).Output()
	pkgs := strings.Split(string(out), "\n")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res []string
	var decoded Package
	for _, pkg := range pkgs {

		if len(pkg) <= 2 {
			continue // invalid pkg name.
		}

		args = []string{"go", "list", "-json", pkg}
		out, err := exec.Command(args[0], args[1:]...).Output()
		err = json.Unmarshal([]byte(out), &decoded)

		if err != nil {
			fmt.Println(string(pkg))
			fmt.Println(string(out))
			fmt.Println(err)
			continue // ignore unmarshal error.
		}

		for _, dep := range decoded.Deps {
			if target == dep {
				res = append(res, decoded.ImportPath)
			}
		}
	}

	fmt.Println(res)
}
