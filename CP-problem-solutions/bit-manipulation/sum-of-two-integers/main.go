//https://leetcode.com/problems/sum-of-two-integers/description/
//Bit manipulation

package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		temp := (a & b) << 1
		a = a ^ b
		b = temp
	}
	return a
}

func main() {
	fmt.Println(getSum(4,5))
	fmt.Println(getSum(9,11))
}
