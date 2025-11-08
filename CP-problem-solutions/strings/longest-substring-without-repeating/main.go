// https://leetcode.com/problems/longest-substring-without-repeating-characters
package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	res := 0
	seen := make([]int, 128)
	start := 0
	for end := 0; end < n; end++ {
		current := s[end]
		if seen[current] > start {
			start = seen[current]
		}
		if end-start+1 > res {
			res = end - start + 1
		}
		seen[current] = end + 1
	}

	return res
}
