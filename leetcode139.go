// +build ignore

package main

import "fmt"

func main() {
	fmt.Print(wordBreak("thisisadog", []string{"this", "thisis", "a", "dog"}))
}

func wordBreak(s string, wordDict []string) bool {
	front := ""
	for i := 0; i < len(s); i++ {
		front += string(s[i])
		judge := false
		for j := 0; j < len(wordDict); j++ {
			if front == wordDict[j] {
				judge = true
				break
			}
		}
		if judge {
			if i == len(s)-1 {
				return true
			}
			rear := s[i+1:]
			result := wordBreak(rear, wordDict)
			if result {
				return result
			}
		}
	}
	return false
}
