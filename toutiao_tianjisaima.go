package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	ours := make([]int, 0)
	yours := make([]int, 0)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		ours = append(ours, tmp)
	}
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		yours = append(yours, tmp)
	}
	res := 0
	for i, j := 0, 0; i < n && j < n; i, j = i+1, j+1 {
		for j < n && yours[i] >= ours[j] {
			j++
		}
		if j < n {
			res++
		}
	}
	fmt.Print(res - (n - res))
}
