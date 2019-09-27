// +build ignore

package main

import "fmt"

func main() {
	//b2 := &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}}
	//b8 := &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}
	//a := &TreeNode{6, b2, b8}
	b1 := &TreeNode{1, nil, nil}
	a := &TreeNode{2, b1, nil}
	c := lowestCommonAncestor(a, a, b1)
	fmt.Print(c.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func search(root, p, q *TreeNode) (bool, *TreeNode) {
	if root == nil {
		return false, nil
	}
	found_left, n1 := search(root.Left, p, q)
	found_right, n2 := search(root.Right, p, q)
	found_root := false
	if n1 != nil {
		return true, n1
	}
	if n2 != nil {
		return true, n2
	}
	if p == root {
		found_root = true
	}
	if q == root {
		found_root = true
	}
	if (found_left && found_right) || (found_root && (found_left || found_right)) {
		return true, root
	}
	return (found_left || found_right || found_root), nil
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, result := search(root, p, q)
	return result
}
