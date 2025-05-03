package main

import (
	"fmt"
	"slices"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes := []rune(s)
	tRunes := []rune(t)

	slices.Sort(sRunes)
	slices.Sort(tRunes)
	
	for i := range sRunes {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
