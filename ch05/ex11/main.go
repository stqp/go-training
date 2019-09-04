package main

func main() {}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)
	end := false

	visitAll = func(items map[string][]string) {

		for item := range items {
			if !seen[item] {
				seen[item] = true

				mm := make(map[string][]string)
				for _, itemm := range m[item] {

					// ※1ステップで循環している場合しか検知しない。
					for _, itemmm := range m[itemm] {
						if item == itemmm {
							end = true
						}
					}

					mm[itemm] = m[itemm]
				}
				visitAll(mm)

				order = append(order, item)
			}
		}
	}
	visitAll(m)
	if end == true {
		order = make([]string, 0)
	}
	return order
}
