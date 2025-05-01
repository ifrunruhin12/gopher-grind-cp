package main

import (
	"fmt"
	"sort"
)

/* Nested for loop O(n^2)

func containsDuplicate(nums []int) bool {
	res := false
	for _, val := range nums {
		cnt := 0
		for _, val2 := range nums {
			if val == val2 {
				cnt++
			}
		}
		if cnt > 1 {
			res = true
			break
		}
	}
	return res
}
*/

//Complexity O(n log n)
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}


func main() {
	s := []int{1, 2, 3, 4, 3}
	res := containsDuplicate(s)
	fmt.Println(res)
}
