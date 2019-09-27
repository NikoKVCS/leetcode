package main

import (
	"fmt"
	"testing"
)

func powerf3(x int, n int) int {
	ans := 1
	for n != 0 {
		ans *= x
		n--
	}
	return ans
}

func divide(dividend int, divisor int) int {
	flag := 1
	if dividend < 0 {
		flag = 0 - flag
		dividend = 0 - dividend
	}
	if divisor < 0 {
		flag = 0 - flag
		divisor = 0 - divisor
	}
	i := 0
	sum := 0
	accelerator := divisor
	acceleratorNum := 1
	for sum < dividend {
		if sum+divisor > dividend {
			break
		}
		if sum+accelerator > dividend {
			accelerator = divisor
			acceleratorNum = 1
		} else if sum+accelerator <= dividend {
			sum = sum + accelerator
			i += acceleratorNum
			accelerator += divisor
			acceleratorNum += 1
		}
	}
	res := i
	if flag < 0 {
		res = 0 - i
	}
	if !(res >= -2147483648 && res <= 2147483647) {
		return 2147483647
	}
	return res
}

func TestLeetcode29(t *testing.T) {
	fmt.Println(powerf3(2, 31))
	fmt.Println(divide(1, 1))
}
