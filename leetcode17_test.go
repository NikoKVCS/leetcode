package main

import (
	"fmt"
	"testing"
)

func letterCombinations(digits string) []string {
	hash := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	result := make([]string, 0)
	for _, v := range digits {
		if len(result) == 0 {
			result = hash[v]
		} else {
			tmp := make([]string, 0)
			for _, item := range result {
				for _, letter := range hash[v] {
					tmp = append(tmp, item+letter)
				}
				result = tmp
			}
		}
	}
	return result
}

func TestLeetcode17(t *testing.T) {
	fmt.Print(letterCombinations("234"))
}
