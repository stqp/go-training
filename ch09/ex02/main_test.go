package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	c := PopCount(0xFFF)
	if c != 12 {
		t.Errorf("invalid : %d", c)
	}
}
