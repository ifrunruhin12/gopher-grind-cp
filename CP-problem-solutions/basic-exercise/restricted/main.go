//https://atcoder.jp/contests/abc063/tasks/abc063_a
//Basic exercise

package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		return
	}
	res := a + b
	if res < 10 {
		fmt.Println(res)
	}else {
		fmt.Println("error")
	}
}
