//https://leetcode.com/problems/counting-bits/
//Bit manipulation

package main

import "fmt"

func dectobin(n int) int {
	bin := fmt.Sprintf("%b", n)
	cnt := 0
	for _, v := range bin {
		if v == '1' {
			cnt++
		}
	}
	return cnt
}

func countBits(n int) []int {
	var s []int
	for i:=0;i<=n;i++{
		count := dectobin(i)
		s = append(s, count)
	}
	return s
}

func main() {
	n := 5
	fmt.Println(countBits(n))
}
