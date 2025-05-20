//https://codeforces.com/problemset/problem/2008/A
package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		total := a + 2*b

		// If total sum is odd, can't split into two equal parts
		if total%2 != 0 {
			fmt.Println("NO")
			continue
		}

		target := total / 2

		// Try to use as many 2s as possible to form target
		twosUsed := min(b, target/2)
		remaining := target - (twosUsed * 2)

		if remaining <= a {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

