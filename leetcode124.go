// +build ignore

package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxx int

func max(nums []int) int {
	best := nums[0]
	for _, item := range nums {
		if item > best {
			best = item
		}
	}
	return best
}
func maxPathSum2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l1 := maxPathSum2(root.Left)
	l2 := maxPathSum2(root.Right)
	bestTree := max([]int{root.Val + l1, root.Val + l2, 0})
	maxx = max([]int{maxx, l1 + l2 + root.Val})
	return bestTree
}
func maxPathSum(root *TreeNode) int {
	maxx = root.Val
	maxPathSum2(root)
	return maxx
}
func main() {
	a := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	//a := &TreeNode{-3, nil, nil}
	//a := &TreeNode{2, &TreeNode{-1, nil, nil}, nil}
	fmt.Print(maxPathSum(a))
}
