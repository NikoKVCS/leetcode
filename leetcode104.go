// +build ignore

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return depth(root, 0)
}
func depth(node *TreeNode, h int) int {
	if node == nil {
		return h
	}
	h++
	if node.Left == nil && node.Right == nil {
		return h
	}
	h1 := depth(node.Left, h)
	h2 := depth(node.Right, h)
	if h1 > h2 {
		return h1
	} else {
		return h2
	}
}
