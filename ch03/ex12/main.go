package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("%v.\n", isAnagram(os.Args[1], os.Args[2]))
}

type anagramTable map[byte]int

var (
	ignored = []byte{
		' ', '\'',
	}
)

func isSameString(s1, s2 string) bool {
	for _, i := range ignored {
		s1 = strings.ReplaceAll(s1, string(i), "")
		s2 = strings.ReplaceAll(s2, string(i), "")
	}
	if strings.ToLower(s1) == strings.ToLower(s2) {
		return true
	}
	return false
}

func isIgnored(c byte) bool {
	for i := 0; i < len(ignored); i++ {
		if ignored[i] == c {
			return true
		}
	}
	return false
}

func mapping(table *anagramTable, s string) {
	s = strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		if isIgnored(s[i]) {
			continue
		}
		(*table)[s[i]]++
	}
}

func isAnagram(s1, s2 string) bool {
	if isSameString(s1, s2) {
		return false
	}

	m := make(anagramTable)
	mapping(&m, s1)
	mapping(&m, s2)
	for _, v := range m {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
