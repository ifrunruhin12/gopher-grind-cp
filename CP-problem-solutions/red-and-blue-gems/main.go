//https://www.codechef.com/problems/REDBLUEGEM
package main

import "fmt"

func maxNum (a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var r, b, p, q int
	fmt.Scan(&r, &b, &p, &q)
	fmt.Println(maxNum(r*p, b*q))
}
