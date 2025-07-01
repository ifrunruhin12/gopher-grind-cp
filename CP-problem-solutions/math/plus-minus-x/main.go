//https://atcoder.jp/contests/abc137/tasks/abc137_a?lang=en
//Math

package main

import "fmt"

func maxOfThree(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func main(){
	var a, b int
	fmt.Scan(&a, &b)
	sum := a + b
	sub := a - b
	mul := a * b
	fmt.Println(maxOfThree(sum, sub, mul))
}
