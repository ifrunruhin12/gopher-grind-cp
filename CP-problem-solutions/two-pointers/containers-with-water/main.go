// https://leetcode.com/problems/container-with-most-water/description/
package main

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		h := min(height[l], height[r])
		w := r - l
		a := h * w
		if a > res {
			res = a
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
