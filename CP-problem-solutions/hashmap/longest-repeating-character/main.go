// https://leetcode.com/problems/longest-repeating-character-replacement/description/
package main

func characterReplacement(s string, k int) int {
	charMap := map[byte]int{}
	maxLen := 0
	maxRepeat := byte(0)
	start := 0
	end := 0
	for end < len(s) {
		char := s[end]
		charMap[char]++
		count := charMap[char]
		if maxRepeat == 0 || charMap[maxRepeat] < count {
			maxRepeat = char
		}
		if end-start+1-charMap[maxRepeat] > k {
			charMap[s[start]]--
			start++
		}
		if maxLen < end-start+1 {
			maxLen = end - start + 1
		}
		end++
	}

	return maxLen
}
