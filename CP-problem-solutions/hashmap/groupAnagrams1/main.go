// https://leetcode.com/problems/group-anagrams/description/
package main

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	cnt := make(map[rune]int)
	for _, ch := range a {
		cnt[ch]++
	}
	for _, ch := range b {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	touch := make([]bool, n)
	var res [][]string

	for i := 0; i < n; i++ {
		if touch[i] {
			continue
		}

		group := []string{strs[i]}
		touch[i] = true

		for j := i + 1; j < n; j++ {
			if !touch[j] && isAnagram(strs[i], strs[j]) {
				group = append(group, strs[j])
				touch[j] = true
			}
		}
		res = append(res, group)
	}

	return res
}
