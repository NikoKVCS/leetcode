package main

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	longest := 0
	length := 0
	hash := make(map[rune]int, 0)
	for i, c := range s {
		if prev, ok := hash[c]; ok {
			length = i - prev
			hash = make(map[rune]int, 0)
			for j := i; j > prev; j-- {
				hash[rune(s[j])] = j
			}
		} else {
			length++
			hash[c] = i
		}

		if length > longest {
			longest = length
		}
	}
	return longest
}

func TestLeetcode3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring(" "))
}
