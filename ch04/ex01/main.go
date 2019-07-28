package main

import (
	"crypto/sha256"
	"fmt"
)

func popCount(x [32]byte) int {
	res := 0
	for i := 0; i < len(x); i++ {
		for j := uint(0); j < 8; j++ {
			res += int((x[i] >> j) & 0x01)
		}

	}
	return res
}

func XorPopCount(x, y [32]byte) int {
	var z [32]byte
	for i := 0; i < 32; i++ {
		z[i] = x[i] ^ y[i]
	}
	return popCount(z)
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%d\n", XorPopCount(c1, c2))
}
