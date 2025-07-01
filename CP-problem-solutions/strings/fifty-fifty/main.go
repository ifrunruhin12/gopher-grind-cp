//https://atcoder.jp/contests/abc132/tasks/abc132_a?lang=en
//String

package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s) != 4 {
		return false
	}

	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	if len(freq) != 2 {
		return false
	}

	for _, count := range freq {
		if count != 2 {
			return false
		}
	}

	return true
}

func main() {
	var s string
	fmt.Scan(&s)

	if isValid(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

