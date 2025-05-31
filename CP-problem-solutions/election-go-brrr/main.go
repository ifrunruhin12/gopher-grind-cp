//https://atcoder.jp/contests/abc366/tasks/abc366_a?lang=en
//Implementation

package main

import "fmt"

func max_num(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	maxNum := max_num(b,c)
	p_vote := a - maxNum
	if maxNum > p_vote {
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}
