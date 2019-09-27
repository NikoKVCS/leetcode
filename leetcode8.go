// +build ignore

package main
import "fmt"

import "strings"
import "math"

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	n := len(str)
	result := 0
	flag := 1
	for i := 0; i<n; i++{
		c := str[i]
		if int(c) >= int('0') && int(c) <= int('9'){
			var bignum int64 = (int64(result) * 10 + int64(c) - int64('0'))*int64(flag)
			if bignum > math.MaxInt32{
				return math.MaxInt32
			} else if bignum < math.MinInt32{
				return math.MinInt32
			} else{
				result = result * 10 + int(c) - int('0')
			}
		} else if c == '-' && 0 == i{
			flag *= -1
		} else if c != '+' || 0 != i{
			break
		}
	}
    return flag*result
}
func main(){
	a := myAtoi("  42")
	fmt.Printf(string(a))
}