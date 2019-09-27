// +build ignore

package main

import "fmt"

func main() {
	a := spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
	fmt.Printf("%v", a)
}

func spiralOrder(matrix [][]int) []int {
	max_i := len(matrix)
	if max_i <= 0 {
		return make([]int, 0)
	}
	min_i := -1
	min_j := -1
	max_j := len(matrix[0])

	i := 0
	j := 0
	dir_ver := 0
	dir_hor := 1
	result := make([]int, 0)
	for i > min_i && i < max_i && j > min_j && j < max_j {
		result = append(result, matrix[i][j])
		if dir_hor > 0 && j+dir_hor >= max_j {
			dir_hor = 0
			dir_ver = 1
			min_i++
		} else if dir_hor < 0 && j+dir_hor <= min_j {
			dir_hor = 0
			dir_ver = -1
			max_i--
		} else if dir_ver < 0 && i+dir_ver <= min_i {
			dir_hor = 1
			dir_ver = 0
			min_j++
		} else if dir_ver > 0 && i+dir_ver >= max_i {
			dir_hor = -1
			dir_ver = 0
			max_j--
		}
		i += dir_ver
		j += dir_hor
	}
	return result
}
