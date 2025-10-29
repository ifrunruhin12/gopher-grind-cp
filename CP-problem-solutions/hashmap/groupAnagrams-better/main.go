//https://leetcode.com/problems/group-anagrams/description/

package main

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, word := range strs {
		var freq [26]int
		for _, ch := range word {
			freq[ch-'a']++
		}
		groups[freq] = append(groups[freq], word)
	}

	res := make([][]string, 0, len(groups))
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}
