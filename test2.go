// +build ignore

package main

import "fmt"

func main() {
	a := []int{1, 3, 4}
	a = a[:len(a)-1]
	fmt.Print(a)
}
