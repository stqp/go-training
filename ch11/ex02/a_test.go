package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		input []int
	}{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{0, 10, 20, 99}},
		{[]int{100, -50, -10, 0, 10}},
	}
	for _, test := range tests {

		intset := IntSet{}
		mapset := MapSet{}

		for _, i := range test.input {
			intset.Add(i)
			mapset.Add(i)
		}

		for i := 0; i < 50; i++ {
			if intset.Has(i) != mapset.Has(i) {
				t.Errorf("intset.Has(%v) = %v ,  mapset.Has(%v) = %v", i, intset.Has(i), i, mapset.Has(i))
			}

		}
	}
}

func TestIntsetUnionWith(t *testing.T) {
	var tests = []struct {
		in1 []int
		in2 []int
	}{
		{
			[]int{1, 2, 3},
			[]int{4, 5},
		},
		{
			[]int{0, 10},
			[]int{20, 999},
		},
		{
			[]int{3, 1, 4, 5, 0, 10, 99},
			[]int{0, 10},
		},
	}
	for _, test := range tests {

		intset1 := IntSet{}
		intset2 := IntSet{}
		mapset1 := MapSet{}
		mapset2 := MapSet{}

		for _, i := range test.in1 {
			intset1.Add(i)
			mapset1.Add(i)
		}
		for _, i := range test.in2 {
			intset2.Add(i)
			mapset2.Add(i)
		}

		intset1.UnionWith(&intset2)
		mapset1.UnionWith(&mapset2)

		for i := -10; i < 1000; i++ {
			if intset1.Has(i) != mapset1.Has(i) {
				t.Errorf("intset.Has(%v) = %v ,  mapset.Has(%v) = %v", i, intset1.Has(i), i, mapset1.Has(i))
			}

		}
	}
}
