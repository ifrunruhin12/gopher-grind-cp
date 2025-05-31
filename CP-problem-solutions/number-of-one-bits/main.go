//https://leetcode.com/problems/number-of-1-bits/description/
//Bit Manipulation

package main

import (
	"fmt"
)

func hammingWeight(n int) int {
	bin := fmt.Sprintf("%b", n)
	cnt := 0
	for _, ch := range bin {
		if ch == '1' {
			cnt++
		}
	}
	return cnt
}

func main() {
	n := 2147483645
	fmt.Println(hammingWeight(n))
}
