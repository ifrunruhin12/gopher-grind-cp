//https://atcoder.jp/contests/abc043/tasks/abc043_a?lang=en

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	res := (n * (n + 1)) / 2
	fmt.Println(res)
}
