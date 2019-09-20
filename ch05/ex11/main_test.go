package main

import (
	"testing"
)

var prereqs1 = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var prereqs2 = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var prereqs3 = map[string][]string{
	"algorithms":           {"data structures"},
	"calculus":             {"linear algebra"},
	"intro to programming": {"programming languages"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func TestMain(t *testing.T) {
	prereqs := prereqs1
	actual, _ := topoSort(prereqs)

	if len(actual) != 13 {
		t.Error("Expected:", len(prereqs), "Actual:", len(actual))
	}

	// preコースとpostコースの前後関係を満たすか、全部ループして調べる。
	for postreqCourse, prereq := range prereqs {
		for _, prereqCourse := range prereq {
			foundPostreq := false
			for _, course := range actual {
				if postreqCourse == course {
					foundPostreq = true
				}
				if prereqCourse == course {
					if foundPostreq == true {
						t.Errorf("Prereqs is not satisfied. %s -> %s", postreqCourse, prereqCourse)
					}
				}
			}
		}
	}

	actual, err := topoSort(prereqs2)
	if err == nil {
		t.Error("failed to find cyclic")
	}

	actual, err = topoSort(prereqs3)
	if err == nil {
		t.Error("failed to find cyclic")
	}

}
