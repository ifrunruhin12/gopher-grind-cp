package main

import "fmt"

func main() {
	var t, a, b, x, y, p, q int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b, &x, &y, &p, &q)
		if (a == x && b == y) || (a == y && b == x) {
			fmt.Println(1)
		} else if (a == p && b == q) || (a == q && b == p) {
			fmt.Println(2)
		} else {
			fmt.Println(0)
		}
	}
}
