package main

import (
	"fmt"
	"testing"
)

func sliceToLinkNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	f := ListNode{
		arr[0],
		nil,
	}
	p := &f
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{arr[i], nil}
		p = p.Next
	}
	return &f
}
func linkNodeToSlice(l *ListNode) []int {
	ret := []int{}
	n := l
	for n != nil {
		ret = append(ret, n.Val)
		n = n.Next
	}
	return ret
}
func TestMergeKLists(t *testing.T) {
	fl := sliceToLinkNode([]int{1, 2, 3})
	sl := sliceToLinkNode([]int{1, 2, 3})
	tl := sliceToLinkNode([]int{1, 2, 3})
	fmt.Println(linkNodeToSlice(mergeKLists([]*ListNode{fl, sl, tl})))
}
