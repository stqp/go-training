package main

import (
	"errors"
)

func main() {}

func keys(m map[string][]string) []string {
	var a []string
	for k := range m {
		a = append(a, k)
	}
	return a
}

func topoSort(m map[string][]string) (order []string, err error) {

	seen := make(map[string]int)
	var visit func(items string)

	visit = func(item string) {
		if seen[item] == 1 {
			err = errors.New("cyclic found")

		} else if seen[item] == 0 {
			seen[item] = 1
			for _, nextItem := range m[item] {
				visit(nextItem)
			}
			seen[item] = 2
			order = append(order, item)
		}
	}

	for item := range m {
		if seen[item] == 0 {
			visit(item)
		}
	}

	return order, err
}
