//https://www.codechef.com/problems/FIT
//Math

package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(n * 10)
	}
}

