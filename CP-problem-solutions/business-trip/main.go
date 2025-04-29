/*
https://codeforces.com/contest/149/problem/A
*/



package main

import (
	"fmt"
	"sort"
)

func main() {
	var k int
	// Read k
	fmt.Scan(&k)
	
	// Create an array of 12 integers
	var water [12]int
	for i := 0; i < 12; i++ {
		fmt.Scan(&water[i])
	}

	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(water[:])))

	sum := 0
	month := 0
	found := false

	// If k is 0, Petya doesn't need to water the plant at all
	if k == 0 {
		fmt.Println(0)
		return
	}

	// Try to accumulate water until we reach or exceed k
	for i := 0; i < 12; i++ {
		sum += water[i]
		month++
		if sum >= k {
			fmt.Println(month)
			found = true
			break
		}
	}

	// If we couldn't reach k, print -1
	if !found {
		fmt.Println(-1)
	}
}
