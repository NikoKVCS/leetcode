// +build ignore

package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return make([][]int, 0)
	}
	result := append(make([][]int, 0), []int{})
	for i := 1; i <= len(nums); i++ {
		result = append(result, C(nums, i)...)
	}
	return result
}

func C(nums []int, n int) [][]int {
	result := make([][]int, 0)
	if n == 1 {
		for _, item := range nums {
			result = append(result, []int{item})
		}
		return result
	}
	for len(nums) >= n {
		a := nums[0]
		nums = append([]int{}, nums[1:]...)
		lists := C(nums, n-1)
		for j, item := range lists {
			lists[j] = append(item, a)
		}
		result = append(result, lists...)
	}
	return result
}
func main() {
	fmt.Print(subsets([]int{1, 2, 3}))
}
