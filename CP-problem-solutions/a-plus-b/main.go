//https://acm.timus.ru/problem.aspx?space=1&num=1000

package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		return
	}
	fmt.Println(a + b)
}

