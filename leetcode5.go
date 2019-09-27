// +build ignore

package main
import "fmt"
import "strings"

func longestPalindrome(s string) string {
    s_new := initString(s)
    n := len(s_new)
    p := make([]int, n)
    mx := 0
    id := 0
    max_id := 0
    max_r := 1
    for i:=1; i<n-1; i=i+1{
        if i< mx{
            p[i] = Min(p[2*id-i], mx-i)
        } else {
            p[i] = 1
        }
        for s_new[i-p[i]] == s_new[i+p[i]] {
            p[i]++
        }
        if mx < i+p[i]{
            mx = i+p[i]
            id = i
        }
        if p[i] >= max_r{
            max_r = p[i]
            max_id = i
        }
    }
    result := s_new[max_id-(max_r-1):max_id+(max_r-1)]
    result = strings.Replace(result,"#","",-1)
    result = strings.Replace(result,"$","",-1)
    result = strings.Replace(result,"&","",-1)
    return result
}
func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
func initString(s string ) string {
    n := len(s)
    s_new := "$#"
    for i:=0;i<n;i=i+1{
        s_new += string(s[i])
        s_new += "#"
    }
    s_new += "&" //保证边界不同
    return s_new
}
func main(){
    s:=longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
    fmt.Printf(s)
    fmt.Printf("\n")
}