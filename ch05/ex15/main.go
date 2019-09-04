package main

func main() {}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	res := vals[0]
	for _, v := range vals {
		if res < v {
			res = v
		}
	}
	return res
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	res := vals[0]
	for _, v := range vals {
		if res > v {
			res = v
		}
	}
	return res
}
