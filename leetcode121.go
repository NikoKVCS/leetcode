// +build ignore

package main

import "fmt"

func main() {
	fmt.Print(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	min := -1
	best := 0
	n := len(prices)
	for i := 0; i < n; i++ {
		if min == -1 || prices[i] < min {
			min = prices[i]
		}
		delta := prices[i] - min
		if delta > best {
			best = delta
		}
	}
	return best
}
