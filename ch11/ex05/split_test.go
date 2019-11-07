package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"aa:bbb:cccc", ":", 3},
		{"::a::b::c::", ":", 9},
	}
	for _, a := range tests {
		words := strings.Split(a.s, a.sep)
		if got := len(words); got != a.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", a.s, a.sep, got, a.want)
		}
	}

}
