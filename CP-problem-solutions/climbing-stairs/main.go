//https://leetcode.com/problems/climbing-stairs/description/

package main

import "fmt"

func climbStairs(n int) int {
    one, two := 1, 1
    for i:=0;i<n-1;i++{
        temp := one
        one = one + two
        two = temp
    }
    return one
}

func main() {
	n := 5
	fmt.Println(climbStairs(n))
}
