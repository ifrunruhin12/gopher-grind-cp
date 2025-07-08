//https://codeforces.com/problemset/problem/313/A
package main

import "fmt"

func maxTwoNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	withoutLast := n / 10
	lastDigit := n % 10
	withoutSecondLast := (n / 100) * 10 + lastDigit
	if n < 0 {
		fmt.Println(maxTwoNum(withoutLast, withoutSecondLast))
	} else {
		fmt.Println(n)
	}
}
