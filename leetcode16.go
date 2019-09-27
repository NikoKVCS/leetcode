// +build ignore

package main

import "sort"

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}
func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(a, b int) bool {
		if nums[a] < nums[b] {
			return true
		}
		return false
	})
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			c := nums[i] + nums[left] + nums[right]
			if abs(c-target) < abs(closest-target) {
				closest = c
			}
			if c == target {
				return target
			} else if c > target {
				right--
			} else {
				left++
			}
		}
	}
	return closest
}
