//https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/description/

//two-pointers

package main

import (
	"fmt"
	"sort"
)

//Nested loop
/*
func countPairs(nums []int, target int) int{
	cnt := 0
	for i:=0;i<len(nums);i++{
		for j:= i + 1;j<len(nums);j++{
			if ((nums[i] + nums[j]) < target) && (i >= 0) && (j > i) {
				cnt++
			}
		}
	}
	return cnt
}
*/

//two-pointers

func countPairs(nums []int, target int) int {
    sort.Ints(nums)
    count := 0
    left, right := 0, len(nums)-1

    for left < right {
        if nums[left]+nums[right] < target {
            count += right - left
            left++
        } else {
            right--
        }
    }

    return count
}

func main() {
	s := []int{3, 4, 6, 2, 1}
	res := countPairs(s, 4)
	fmt.Println(res)
}
