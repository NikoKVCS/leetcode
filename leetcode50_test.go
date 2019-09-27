package main

import (
	"fmt"
	"testing"
)

func myPow(x float64, n int) float64 {
	flag := 1
	if n < 0 {
		flag = -1
		n *= flag
	}
	res := 1.0
	for i := n; i != 0; i /= 2 {
		if i%2 != 0 {
			res *= x
		}
		x *= x
	}
	if flag == -1 {
		res = 1 / res
	}
	return res
}

func TestLeetcode50(t *testing.T) {
	fmt.Println(searchRange([]int{1}, 1))
}
