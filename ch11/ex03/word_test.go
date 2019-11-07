package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNotPalindrome(rng *rand.Rand) string {
	for {
		n := rng.Intn(25)
		if n < 2 {
			continue
		}

		runes := make([]rune, n)
		for i := 0; i < n; i++ {
			for {
				r := rune(rng.Intn(0x1000))
				if unicode.IsLetter(r) {
					runes[i] = r
					break
				}
			}
		}

		for i := 0; i < len(runes)/2; i++ {
			if runes[i] != runes[len(runes)-1-i] {
				return string(runes)
			}
		}
	}
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 10; i++ {
		p := randomNotPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}

}
