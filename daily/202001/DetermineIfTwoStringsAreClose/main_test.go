package main

import "testing"

func TestCloseStrings(t *testing.T) {
	table := []struct {
		word1  string
		word2  string
		expect bool
	}{
		{"abcc", "abcac", false},
		{"aaaaaa", "bbbbbb", false},
		{"abcabc", "cbacba", true},
		{"uau", "ssx", false},
	}
	for _, d := range table {
		a := closeStrings(d.word1, d.word2)
		if d.expect != a {
			t.Errorf("Error: Word1=%s, Word2=%s expected value is %v, but actual is %v", d.word1, d.word2, d.expect, t)
		}
	}
}
