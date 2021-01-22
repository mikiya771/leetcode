package main

import "testing"

func TestIsOneEditDistance(t *testing.T) {
	table := []struct {
		word1  string
		word2  string
		expect bool
	}{
		{"abc", "abc", false},
		{"abc", "bcd", false},
		{"abc", "abd", true},
		{"ab", "abc", true},
		{"abc", "ab", true},
		{"ab", "adb", true},
		{"db", "adb", true},
		{"db", "adbc", false},
		{"abc", "ef", false},
		{"ef", "abc", false},
	}
	for _, d := range table {
		a := isOneEditDistance(d.word1, d.word2)
		if d.expect != a {
			t.Errorf("Error: Word1=%s, Word2=%s expected value is %v, but actual is %v", d.word1, d.word2, d.expect, t)
		}
	}
}
