package main

import "math"

func kLengthApart(nums []int, k int) bool {
	span := calcSmallestSpan(nums)
	return span >= k
}

func calcSmallestSpan(nums []int) int {
	min := math.MaxInt32
	count := 0
	isCount := false
	for _, num := range nums {
		if num == 1 {
			if min > count && isCount {
				min = count
			}
			count = 0
			isCount = true
			continue
		}
		if isCount {
			count++
		}
	}
	return min
}
