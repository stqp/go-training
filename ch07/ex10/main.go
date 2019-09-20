package main

import "sort"

func main() {}

func IsPalindrome(s sort.Interface) bool {
	l := s.Len()
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
