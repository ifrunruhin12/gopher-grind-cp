//https://atcoder.jp/contests/abc105/tasks/abc105_a?lang=en

package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	if N%K == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}

