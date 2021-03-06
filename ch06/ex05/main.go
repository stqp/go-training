package intset

import (
	"bytes"
	"fmt"
)

// IntSet Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

const (
	uintBits = 32 << (^uint(0) >> 63)
)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintBits, uint(x%uintBits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintBits, uint(x%uintBits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll is
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith is
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWith is
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] = (s.words[i] ^ t.words[i]) & s.words[i]
		}
	}
}

// SymmetricDifference is
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintBits; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintBits*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) all() (res []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintBits; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, uintBits*i+j)
			}
		}
	}
	return res
}

// Len is
func (s *IntSet) Len() int {
	res := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintBits; j++ {
			if word&(1<<uint(j)) != 0 {
				res++
			}
		}
	}
	return res
}

// Remove is
func (s *IntSet) Remove(x int) {
	word, bit := x/uintBits, uint(x%uintBits)
	if word > len(s.words) {
		return
	}
	if s.words[word]&(1<<bit) != 0 {
		s.words[word] &= ^(uint(1) << bit)
	}
	s.words[word] |= 0 << bit
}

// Clear is
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy is
func (s *IntSet) Copy() *IntSet {
	new := &IntSet{}

	for _, v := range s.all() {
		new.Add(v)
	}
	return new
}

// Elems is
func (s *IntSet) Elems() []int {
	return s.all()
}
