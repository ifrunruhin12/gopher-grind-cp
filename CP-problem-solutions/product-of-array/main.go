//https://leetcode.com/problems/product-of-array-except-self/description/

package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}

	prefix := 1

	for j := range res {
		res[j] = prefix
		prefix *= nums[j]
	}

	postfix := 1

	for k := range res {
		res[(len(nums)-1)-k] *= postfix
		postfix *= nums[(len(nums)-1)-k]
	}
	return res
}

func main() {
	slc := []int{1,2,3,4}
	res := productExceptSelf(slc)
	fmt.Println(res)
}
