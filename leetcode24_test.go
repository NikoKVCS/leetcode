package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	A := head
	B := head.Next
	pre := &ListNode{0, A}
	node := pre
	for A != nil && B != nil {
		store := B.Next
		B.Next = A
		A.Next = store
		pre.Next = B
		pre = A
		A = A.Next
		if B.Next != nil && B.Next.Next != nil {
			B = B.Next.Next.Next
		}
	}
	return node.Next
}
func TestLeetcode24(t *testing.T) {
	fmt.Println(swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
}
