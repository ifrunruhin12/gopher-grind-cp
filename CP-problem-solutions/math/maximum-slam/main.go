//https://www.codechef.com/START190D/problems/MAXSLAM

package main

import "fmt"

func ceil(a, b int) int {
	return (a + b - 1) / b
}

func main() {
	var n int
	fmt.Scan(&n)
	temp := 25 - n
	fmt.Println(ceil(temp, 4))
}

