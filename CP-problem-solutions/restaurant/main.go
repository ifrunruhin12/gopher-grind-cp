//https://atcoder.jp/contests/abc055/tasks/abc055_a?lang=en

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	res := (n * 800) - ((n/15) * 200)
	fmt.Println(res)
}
