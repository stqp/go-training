package main

func main() {}

func rotate(s []int, by int) []int {
	base := -by % len(s)
	if base < 0 {
		base += len(s)
	}
	return append(s[base:], s[:base]...)
}

func RotateRight(s []int, by int) []int {
	return rotate(s, by)
}

func RotateLeft(s []int, by int) []int {
	return rotate(s, -by)
}

func equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
