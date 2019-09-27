// +build ignore

package main

func reverse(x int) int {
    t := x
    y := 0
    for 0 != t/10 || 0 != t%10 {
        y = y*10 + t%10
        t = t/10
    }
    if powerf3(2,31)-1 < y || 0-powerf3(2,31) > y {
        y = 0
    }
    return y
}
func powerf3(x int, n int) int {
    ans := 1
    for n != 0 {
        ans *= x
        n--
    }
    return ans
}