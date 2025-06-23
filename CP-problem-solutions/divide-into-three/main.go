//https://www.codechef.com/problems/DIVIDE3

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%3 == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
