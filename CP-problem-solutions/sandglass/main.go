//https://atcoder.jp/contests/abc072/tasks/abc072_a?lang=en

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a > b {
		fmt.Println(a-b)
	}else {
		fmt.Println(0)
	}
}
