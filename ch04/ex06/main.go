package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	runes := []rune("こん　にちは　　世界")
	encoded := encodeRunes(runes)
	out := replaceUnicodeSpace(encoded)
	fmt.Println(string(out))
}

func encodeRunes(runes []rune) []byte {
	bytes := make([]byte, 300)
	i := 0
	for _, r := range runes {
		t := utf8.EncodeRune(bytes[i:], r)
		i += t
	}
	return bytes[:i]
}

func replaceUnicodeSpace(bytes []byte) []byte {
	out := bytes[:0]
	for i := 0; i < len(bytes); i++ {
		// AND演算について遅延評価されるのでOut of Indexエラーにならない。すばらしい。
		if bytes[i] == 0xE3 && bytes[i+1] == 0x80 && bytes[i+2] == 0x80 {
			out = append(out, 0x20)
			i += 2
		} else {
			out = append(out, bytes[i])
		}
	}
	return out
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
