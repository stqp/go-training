package main

import "sort"

func main() {}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
aaaaa
func topoSort(m map[string][]string) []string {
	var order []string

	var visit func(item string) []string
	visit = func(item string) []string {
		order = append(order, item)
		return m[item]
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	breadthFirst(visit, keys)
	return order
}
