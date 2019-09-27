// +build ignore

package main

import "fmt"

func multiply(num1 string, num2 string) string {
	n := len(num1)
	num_final := make([]int, 1)
	for i := 0; i < n; i++ {
		pow := n - 1 - i
		// 乘法器
		mul1 := int(num1[i] - '0')
		num_temp := make([]int, pow+len(num2)+1)
		f := pow
		for j := len(num2) - 1; j >= 0; j-- {
			mul2 := int(num2[j] - '0')
			num_temp[f] += mul1 * mul2
			num_temp[f+1] += num_temp[f] / 10
			num_temp[f] %= 10
			f++
		}
		for q := 0; q < len(num_temp); q++ {
			if q >= len(num_final)-1 {
				num_final = append(num_final, 0)
			}
			num_final[q] += num_temp[q]
			num_final[q+1] += num_final[q] / 10
			num_final[q] %= 10
		}
	}
	for q := 0; q < len(num_final)-1; q++ {
		num_final[q+1] += num_final[q] / 10
		num_final[q] %= 10
	}
	result := ""
	for i := len(num_final) - 1; i >= 0; i-- {
		if result == "" && num_final[i] == 0 {
			continue
		}
		result = result + string(num_final[i]+'0')
	}
	if result == "" {
		return "0"
	}
	return result
}
func main() {
	num1 := "2925"
	num2 := "4787"
	fmt.Print(multiply(num1, num2))
}
