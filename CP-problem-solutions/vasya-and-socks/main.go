//https://codeforces.com/problemset/problem/460/A

package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	days := 0

	for n > 0 {
		days++    // one more day passes
		n--       // Vasya uses one sock

		if days%m == 0 {
			n++ // mom gives him one sock in the evening
		}
	}

	fmt.Println(days)
}

