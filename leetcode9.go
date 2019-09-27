// +build ignore

package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := 0
	t := x
	for 0 != x/10 || 0 != x%10 {
		y = y*10 + x%10
		x = x / 10
	}
	if t == y {
		return true
	}
	return false
}

func main() {

}
