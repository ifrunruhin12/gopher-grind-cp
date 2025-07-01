//https://leetcode.com/problems/two-sum/description/
//two-pointers

package main

import (
	"fmt"
	"sort"
)

func indexOfOnce(nums []int, target int, skip int) int {
	for i, v := range nums {
		if v == target && i != skip {
			return i
		}
	}
	return -1
}

// Reusable TwoSum with slice copy + two-pointer on sorted + index recovery
func twoSum(nums []int, target int) []int {
	// Copy and sort the input
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	// Two-pointer to find the two values
	left, right := 0, len(sorted)-1
	for left < right {
		sum := sorted[left] + sorted[right]
		if sum == target {
			val1 := sorted[left]
			val2 := sorted[right]
			// Find original indices in the unsorted array
			i1 := indexOfOnce(nums, val1, -1)
			i2 := indexOfOnce(nums, val2, i1)
			return []int{i1, i2}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main(){
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
