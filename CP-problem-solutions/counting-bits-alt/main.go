//https://leetcode.com/problems/counting-bits/

package main

import "fmt"

func countBits(n int) []int {
	dp := make([]int, n+1)
	offset := 1

	for i:=1;i<=n;i++{
		if offset * 2 == i {
			offset = i
		}
		dp[i] = 1 + dp[i - offset]
	}
	return dp
}

func main() {
	n := 5
	fmt.Println(countBits(n))
}
