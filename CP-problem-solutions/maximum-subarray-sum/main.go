//https://leetcode.com/problems/maximum-subarray/description/
//Divide and Conquere
//DP

package main

import "fmt"

func maxOfTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	maxSub := nums[0]
	curSum := 0

	for _, v := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += v
		maxSub = maxOfTwo(curSum, maxSub)
	}
	return maxSub
}

func main() {
	slc := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(slc))
}
