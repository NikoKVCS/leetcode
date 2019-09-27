// +build ignore

package main

import "fmt"


func main(){
	a := 0
	a ^= a
	fmt.Printf(string(a))
}