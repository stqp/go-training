package main

import "fmt"

type myarray [8]byte

func main() {
	a := myarray{1, 2, 3, 4, 5, 6, 7, 8}
	reverse(&a)
	fmt.Println(a)
}

func reverse(s *myarray) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equals(a *myarray, b *myarray) bool {
	if len(*a) != len(*b) {
		return false
	}
	for i := 0; i < len(*a); i++ {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}
	return true
}
