//https://codeforces.com/problemset/problem/546/A
package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	total := a * c * (c + 1) / 2
	ans := total - b

	if ans > 0 {
		fmt.Println(ans)
	} else {
		fmt.Println(0)
	}
}

