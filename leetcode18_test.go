package main

import (
	"fmt"
	"sort"
	"testing"
)

func kSum(num, target int, list []int) [][]int {
	result := make([][]int, 0)
	if num > len(list) {
		return result
	}
	if num == 2 {
		left := 0
		right := len(list) - 1
		for left < right {
			if list[left]+list[right] > target {
				right--
			} else if list[left]+list[right] < target {
				left++
			} else {
				item := []int{list[left], list[right]}
				result = append(result, item)
				left++
				right--
				for list[left] == list[left-1] && left+1 < len(list) {
					left++
				}
				for list[right] == list[right+1] && right-1 > 0 {
					right--
				}
			}
		}
	} else {
		for i, v := range list {
			if i-1 >= 0 && v == list[i-1] {
				continue
			}
			if i > len(list)-num {
				break
			}
			res := kSum(num-1, target-v, list[i+1:])
			for _, item := range res {
				result = append(result, append(item, v))
			}
		}
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	return kSum(4, target, nums)
}

func TestLeetcode18(t *testing.T) {
	fmt.Println(fourSum([]int{-1, -5, -5, -3, 2, 5, 0, 4}, -7))
}
