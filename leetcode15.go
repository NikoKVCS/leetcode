// +build ignore

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Slice(nums, func(a, b int) bool {
		if nums[a] < nums[b] {
			return true
		}
		return false
	})
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 保证不重复
		}
		right := len(nums) - 1
		left := i + 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if 0 == sum {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for nums[left] == nums[left-1] && left < right {
					left++
				}
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
func main() {
	var s []int
	fmt.Print(append(s, 3))
	threeSum([]int{3, 0, -2, -1, 1, 2})
}
