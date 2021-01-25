package main

import (
	"math"
	"testing"
)

func TestCalcLongestSpan(t *testing.T) {
	table := []struct {
		name   string
		nums   []int
		expect int
	}{
		{"all1", []int{1, 1, 1, 1}, 0},
		{"all0", []int{1, 0, 0, 0}, math.MaxInt32},
		{"ordinary", []int{1, 0, 0, 1, 0, 1}, 1},
	}
	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			act := calcSmallestSpan(e.nums)
			if act != e.expect {
				t.Errorf("Error: smallest span of %v is %d, but function returns as %d", e.nums, e.expect, act)
			}
		})
	}
}
