/*
https://codeforces.com/contest/158/problem/B
*/
package main

import (
	"fmt"
)

func minTaxis(n int, groups []int) int {
	// Step 1: Count the occurrences of each group size (1, 2, 3, 4)
	count := [5]int{} // Index 0 is unused since group sizes are from 1 to 4
	for _, size := range groups {
		count[size]++
	}

	// Step 2: Start allocating taxis optimally
	taxis := count[4] // Every group of size 4 gets a separate taxi

	// Step 3: Pair groups of size 3 with size 1 if possible
	taxis += count[3] // Each group of 3 initially gets a separate taxi
	count[1] -= count[3] // Reduce the available 1-sized groups
	if count[1] < 0 {
		count[1] = 0
	}

	// Step 4: Pair groups of size 2 together
	taxis += count[2] / 2 // Each two size-2 groups share a taxi
	count[2] %= 2 // If there's one left, it needs to be handled

	// Step 5: If there's one group of size 2 left, try to pair it with up to two size-1 groups
	if count[2] > 0 {
		taxis++ // One taxi needed for this size-2 group
		count[1] -= 2 // Try to pair it with up to two size-1 groups
		if count[1] < 0 {
			count[1] = 0
		}
	}

	// Step 6: If there are size-1 groups left, fit them in taxis (4 per taxi)
	taxis += count[1] / 4
	if count[1]%4 > 0 {
		taxis++
	}

	return taxis
}

func main() {
	var n int
	fmt.Scan(&n)
	groups := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&groups[i])
	}
	fmt.Println(minTaxis(n, groups))
}
