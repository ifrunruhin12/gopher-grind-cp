//https://codeforces.com/problemset/problem/2044/A

package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(n - 1)
	}
}
