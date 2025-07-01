//https://leetcode.com/problems/3sum/solutions/5055810/video-two-pointer-solution/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicates
		}

		j, k := i+1, len(nums)-1

		for j < k {
			total := nums[i] + nums[j] + nums[k]

			if total > 0 {
				k--
			} else if total < 0 {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++

				// skip duplicates for the second number
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}

