package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}

	notLetters := []rune{
		',',
		' ',
		'.',
	}

	ans := []rune{}
	for {
		if len(runes) == 0 {
			break
		}
		if rng.Intn(0x1000)%2 == 0 {
			ans = append(ans, runes[0])
			runes = runes[1:]
		} else { //たぶん実行されるはず
			a := rng.Intn(len(notLetters))
			ans = append(ans, notLetters[a])
		}
	}
	return string(ans)
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
