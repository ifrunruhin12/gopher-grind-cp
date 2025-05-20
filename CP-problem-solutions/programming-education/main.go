//https://atcoder.jp/contests/abc112/tasks/abc112_a?lang=en

package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	if t == 1 {
		fmt.Println("Hello World")
	}else {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Println(a + b)
	}
}
