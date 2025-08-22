//https://www.codechef.com/practice/course/1to2stars/LP1TO201/problems/MAX_DIFF

package main

import (
	"fmt"
)

func main() {
	var t, n, s int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &s)
		if s <= n {
			fmt.Println(s - 0)
		} else {
			diff := s - n
			res := n - diff
			fmt.Println(res)
		}
	}
}
