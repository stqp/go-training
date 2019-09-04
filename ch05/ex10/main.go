package main

func main() {}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)

	visitAll = func(items map[string][]string) {
		for item := range items {
			if !seen[item] {
				seen[item] = true

				mm := make(map[string][]string)
				for _, item := range m[item] {
					mm[item] = m[item]
				}
				visitAll(mm)
				order = append(order, item)
			}
		}
	}
	visitAll(m)
	return order
}
