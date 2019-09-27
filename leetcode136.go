// +build ignore

package main

import "fmt"

func main() {
	fmt.Print(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	hash := make(map[int]int, 0)
	for _, item := range nums {
		if _, ok := hash[item]; ok {
			delete(hash, item)
		} else {
			hash[item] = item
		}
	}
	for k, _ := range hash {
		return k
	}
	return 0
}
