//https://atcoder.jp/contests/abc061/tasks/abc061_a?lang=en
//Math

package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scanf("%d %d %d", &a, &b, &c)
	if err != nil {
		return
	}
	if c >= a && c <= b {
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}
