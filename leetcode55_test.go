package main

import (
	"fmt"
	"testing"
)

func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i, v := range nums {
		if dp[i] == 0 {
			continue
		}
		for j := i; j <= i+v && j < len(nums); j++ {
			dp[j] = 1
		}
	}
	if dp[len(nums)-1] == 1 {
		return true
	} else {
		return false
	}
}

func TestLeetcode55(t *testing.T) {
	fmt.Println(searchRange([]int{1}, 1))
}
