// +build ignore

package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)
	for 
}

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	list := make([]bool, len(s)+1)
	for _, v := range wordDict {
		m[v] = true
	}
	list[0] = true
	l := len(s)
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if list[j] && m[s[j:i]] {
				list[i] = true
				break
			}
		}
	}
	return list[l]
}
