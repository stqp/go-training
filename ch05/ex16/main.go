package main

func main() {}

func join(sep string, xs ...string) string {
	switch len(xs) {
	case 0:
		return ""
	case 1:
		return xs[0]
	}
	res := xs[0]
	for _, x := range xs[1:] {
		res += sep + x
	}
	return res
}
