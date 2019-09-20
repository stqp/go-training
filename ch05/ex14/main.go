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

func topoSort(m map[string][]string) []string {
	var order []string
	var visit func(items string) []string

	visit = func(item string) []string {
		order = append(order, item)
		delete(m, item)
		var res []string
		for _, nextItem := range m[item] {
			for j := range m {
				for _, v := range m[j] {
					if v == nextItem {
						res = append(res, nextItem)
					}
				}
			}
		}

		return res
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	breadthFirst(visit, keys)

	return order
}
