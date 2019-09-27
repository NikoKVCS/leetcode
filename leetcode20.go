// +build ignore

package main

import "fmt"

func isValid(s string) bool {
	hash := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	stack := make([]rune, len(s))
	top := -1
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			top++
			stack[top] = c
		} else if c == ')' || c == '}' || c == ']' {
			if top < 0 {
				return false
			}
			cpr := stack[top]
			top--
			if hash[cpr] != c {
				return false
			}
		}
	}
	if top >= 0 {
		return false
	}
	return true
}
func main() {
	fmt.Print(isValid("[]()"))
}
