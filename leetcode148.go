// +build ignore

package main

import "fmt"

func main() {
	d := &ListNode{3, nil}
	c := &ListNode{1, d}
	b := &ListNode{2, c}
	a := &ListNode{4, b}
	q := sortList(a)
	fmt.Print(q)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeLists(left *ListNode, right *ListNode) *ListNode {
	result := &ListNode{0, nil}
	p := left
	q := right
	c := result

	for p != nil || q != nil {
		if (p != nil && q != nil && p.Val < q.Val) || (p != nil && q == nil) {
			c.Next = p
			p = p.Next
			c = c.Next
		} else {
			c.Next = q
			q = q.Next
			c = c.Next
		}
	}
	return result.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find the split point
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast != nil && fast.Next != nil {
			slow = slow.Next
		}
	}
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)

	return mergeLists(left, right)
}
