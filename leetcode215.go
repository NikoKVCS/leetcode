// +build ignore

package main

import "fmt"

func main() {
	a := []int{2, 1, 6, 3, 6, 2, 3, 7, 6, 5}
	fmt.Print(findKthLargest(a, 5))
}
func findKthLargest(nums []int, k int) int {
	heapBuild(nums, k)
	n := len(nums)
	for i := k - 1; i < n; i++ {
		if nums[0] < nums[i] {
			nums[0], nums[i] = nums[i], nums[0]
			heapBuild(nums, k)
		}
	}
	heapSort(nums, k)
	return nums[k-1]
}

func heapBuild(nums []int, k int) {
	for root := (k - 1) / 2; root >= 0; root-- {
		minheap(root, k, nums)
	}
}

func heapSort(nums []int, k int) {
	for end := k - 1; end >= 0; end-- {
		if nums[0] < nums[end] {
			nums[0], nums[end] = nums[end], nums[0]
			minheap(0, end-1, nums)
		}
	}
}

func minheap(root int, end int, nums []int) {
	for {
		child := root*2 + 1
		if child >= end {
			break
		}
		if child+1 < end && nums[child+1] < nums[child] { // 是小于 不是小于等于
			child = child + 1
		}
		if nums[child] < nums[root] {
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		} else {
			break
		}
	}
}
