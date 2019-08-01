package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	runes := []rune("こんにちは")
	encoded := encodeRunes(runes)
	reversed := reverseRunes(encoded)
	fmt.Printf("%s\n", string(reversed))
}

func encodeRunes(runes []rune) []byte {
	bytes := make([]byte, 300)
	i := 0
	for _, r := range runes {
		i += utf8.EncodeRune(bytes[i:], r)
	}
	return bytes[:i]
}

func reverseRunes(bytes []byte) []byte {
	i := 0
	for {
		size := reverseRune(bytes[i:])
		if size == 0 {
			break
		}
		i += size
	}
	reverse(bytes)
	return bytes
}

func reverseRune(bytes []byte) int {
	_, size := utf8.DecodeRune(bytes)
	reverse(bytes[:size])
	return size
}

func reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func equals(a []rune, b []rune) bool {
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
