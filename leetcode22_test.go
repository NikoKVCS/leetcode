package main

import (
	"fmt"
	"testing"
)

func statusMachine(n, left, right int, prefix string) []string {
	res := make([]string, 0)
	if left+right == 0 {
		return append(res, prefix)
	}
	status := (n - left) - (n - right)
	if status == 0 && left-1 > 0 {
		res = statusMachine(n, left-1, right, prefix+"(")
	} else if status >= 0 {
		if left-1 >= 0 {
			res = statusMachine(n, left-1, right, prefix+"(")
		}
		if right-1 >= 0 {
			res = append(res, statusMachine(n, left, right-1, prefix+")")...)
		}
	}
	return res
}

func generateParenthesis(n int) []string {
	return statusMachine(n, n, n, "")
}

func TestLeetcode22(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}
