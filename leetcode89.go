// +build ignore

package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	}
	list := grayCode(n - 1)
	list2 := make([]int, 0)
	pow := 1
	for i := 0; i < n-1; i++ {
		pow *= 2
	}
	for i := len(list) - 1; i >= 0; i-- {
		list2 = append(list2, list[i]+pow)
	}
	list = append(list, list2...)
	return list
}

func main() {
	fmt.Print(grayCode(3))
}
