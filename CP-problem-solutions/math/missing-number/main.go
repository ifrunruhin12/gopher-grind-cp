// https://leetcode.com/problems/missing-number/
//Math
//Binary search
//Bit manipulation

package main

import (
	"fmt"
)

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func missingNumber(nums []int) int {
	total := len(nums) * (len(nums) + 1) / 2
	return total - sum(nums)
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
