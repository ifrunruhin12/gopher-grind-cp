package main

import "fmt"

func main() {
	var t, n, a, b, c int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &a, &b, &c)
		if (a + c) <= b {
			dish := a + c
			if dish >= n {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			dish := b
			if dish >= n {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
