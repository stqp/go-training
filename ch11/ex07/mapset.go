package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapSet struct {
	words map[int]bool
}

func (s *MapSet) init() {
	if s.words == nil {
		s.words = make(map[int]bool)
	}
}

func (s *MapSet) Has(x int) bool {
	s.init()
	return s.words[x]
}

func (s *MapSet) Add(x int) {
	s.init()
	s.words[x] = true
}

func (s *MapSet) UnionWith(t *MapSet) {
	s.init()
	for k, v := range t.words {
		s.words[k] = v
	}
}

func (s *MapSet) String() string {
	s.init()
	a := []int{}
	for k, v := range s.words {
		if v {
			a = append(a, k)
		}
	}

	sort.Ints(a)

	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range a {
		if i != 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte('}')
	return buf.String()
}
