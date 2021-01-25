package main

import (
	"errors"
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := lists
	dummy := ListNode{0, nil}
	prev := &dummy
	for {
		n, p, m, e := pickNumber(l)
		if e != nil {
			break
		}
		l[p] = m.Next
		fmt.Println(l)
		prev.Next = &ListNode{n, nil}
		prev = prev.Next
	}
	return dummy.Next
}
func pickNumber(lists []*ListNode) (int, int, *ListNode, error) {
	var minNode *ListNode
	min := math.MaxInt32
	status := false
	pos := 0
	for i, ln := range lists {
		if ln == nil {
			continue
		}
		status = true
		if min > ln.Val {
			min = ln.Val
			pos = i
			minNode = ln
		}
	}
	if !status {
		return 0, 0, nil, errors.New("no valid link")
	}
	return min, pos, minNode, nil
}
