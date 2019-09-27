package main

import (
	"fmt"
	"sort"
	"testing"
)

func nextPermutation(nums []int) {
	if len(nums) <= 0 {
		return
	}
	found := false
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			j := i - 1
			k := i
			for ; k < len(nums)-1; k++ {
				if nums[k] > nums[j] && nums[k+1] < nums[j] {
					break
				}
			}
			tmp := nums[j]
			nums[j] = nums[k]
			nums[k] = tmp

			found = true
			break
		}
	}
	if !found {
		sort.Slice(nums, func(a, b int) bool {
			return a > b
		})
	}
}

func TestLeetcode31(t *testing.T) {
	val := []int{1, 3, 2}
	nextPermutation(val)
	fmt.Println(val)
}
