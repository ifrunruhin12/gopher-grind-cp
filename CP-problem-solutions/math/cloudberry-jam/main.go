//https://codeforces.com/problemset/problem/2086/A
//Math

package main

import (
	"fmt"
)

func main() {
	var t, n int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		fmt.Println(n * 2)
	}
}

