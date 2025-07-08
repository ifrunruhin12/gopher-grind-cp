//https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
package main

import "fmt"

func countDigits(num int) int {
	count := 0
	for num > 0 {
		count++
		num = num / 10
	}
	return count
}

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		if countDigits(num)%2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}
