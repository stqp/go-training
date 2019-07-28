package main

import "unicode/utf8"

func main() {

	runes := []rune("こんにちは　　世界")
	var bytes []byte
	utf8.EncodeRune(bytes, runes)
}

/*
func compressUtf8Spaces(bytes []byte) []byte {
	unicode.IsSpace(bytes[0])
}*/

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func equals(a []string, b []string) bool {
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
