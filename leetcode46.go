// +build ignore

package main

import "fmt"

func main() {
	fmt.Printf("%v", permute([]int{1, 2, 3}))
}
func backtrace(nums []int, res [][]int, i int) [][]int {
	if i == len(nums) {
		tmp2 := make([]int, len(nums))
		copy(tmp2, nums)
		tmp := append(res, tmp2)
		return tmp
	}
	for j := i; j < len(nums); j++ {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		res = backtrace(nums, res, i+1)
		nums[j] = nums[i]
		nums[i] = temp
	}
	return res
}
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	return backtrace(nums, result, 0)
}
