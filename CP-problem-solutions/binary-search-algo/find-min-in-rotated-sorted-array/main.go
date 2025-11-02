// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
package main

func findMin(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
