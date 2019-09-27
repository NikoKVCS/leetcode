// +build ignore

package main

import "fmt"

func main() {
	nums := []int{7, 2, 6, 4, 1, 9, 4}
	QuickSort(nums)
	fmt.Print(nums)
}

func QuickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	key := nums[0]
	i := 0
	j := n - 1
	for i != j {
		for ; j >= i; j-- {
			if nums[j] <= key {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
		for ; i < j; i++ {
			if nums[i] > key {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	QuickSort(nums[:i])
	if i < n {
		QuickSort(nums[:i+1])
	}
}
