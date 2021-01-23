package main

import "testing"

func TestCloseStrings(t *testing.T) {
	table := []struct {
		input  [][]int
		output [][]int
	}{
		{
			[][]int{
				{0, 0, 0},
				{1, 2, 3},
				{4, 6, 1},
			},
			[][]int{
				{0, 0, 0},
				{1, 1, 3},
				{4, 6, 2},
			},
		},
		{
			[][]int{
				{3, 3, 1, 1},
				{2, 2, 1, 2},
				{1, 1, 1, 2},
			},
			[][]int{
				{1, 1, 1, 1},
				{1, 2, 2, 2},
				{1, 2, 3, 3},
			},
		},
	}
	for _, d := range table {
		a := diagonalSort(d.input)
		if !equal(d.output, a) {
			t.Errorf("Error: input %v, expect %v, actual %v", d.input, d.output, a)
		}
	}
}
func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, la := range a {
		lb := b[i]
		if len(la) != len(lb) {
			return false
		}
		for j, ea := range la {
			if ea != lb[j] {
				return false
			}
		}
	}
	return true
}
