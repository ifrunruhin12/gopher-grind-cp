//https://leetcode.com/problems/valid-palindrome/description/
//string
//two-pointer

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	// Two-pointer approach after filtering
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric on the left
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		}
		// Skip non-alphanumeric on the right
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		}

		// Compare lowercased letters
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s)) // Output: true
}

