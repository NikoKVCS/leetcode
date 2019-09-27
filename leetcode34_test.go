package main

import (
	"fmt"
	"testing"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := 0
	right := len(nums) - 1
	begin := -1
	for left <= right {
		if nums[left] == target {
			begin = left
			break
		} else if nums[right] == target {
			begin = right
			break
		} else {
			mid := (left + right) / 2
			if mid == left || mid == right {
				break
			}
			if nums[mid] > target {
				right = mid
			} else {
				left = mid
			}
		}
	}
	if begin == -1 {
		return []int{-1, -1}
	}
	for begin-1 > 0 && nums[begin-1] == target {
		begin--
	}
	end := begin
	for end+1 < len(nums) && nums[end+1] == target {
		end++
	}
	return []int{begin, end}
}

func TestLeetcode34(t *testing.T) {
	fmt.Println(searchRange([]int{1}, 1))
}
