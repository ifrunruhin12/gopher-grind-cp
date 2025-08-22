// https://www.codechef.com/practice/course/1to2stars/LP1TO201/problems/PROBCAT
package main

import "fmt"

func main() {
	var t, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		if n >= 1 && n < 100 {
			fmt.Println("Easy")
		} else if n >= 100 && n < 200 {
			fmt.Println("Medium")
		} else {
			fmt.Println("Hard")
		}
	}
}
