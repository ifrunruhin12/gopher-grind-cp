package main

import "fmt"

func minOfThree(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func main() {
	var t, a, b, c int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b, &c)
		sum := a + b + c
		fmt.Println(sum - minOfThree(a, b, c))
	}
}
