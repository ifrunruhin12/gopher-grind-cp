/*
https://codeforces.com/contest/469/problem/A
*/

package main

import "fmt"

func solve(n int, list1, list2 []int) string {
	newLst := append(list1, list2...)
	cnt := 0
	seen := make(map[int]bool)
	for _, elem := range newLst {
		seen[elem] = true
	}
	for i := 1; i <= n; i++ {
	    if seen[i] {
			cnt++
		}
	}
	if cnt  == n {
		return "I become the guy."
	}
	return "Oh, my keyboard!"
}

func main() {
	var n int
	fmt.Scan(&n)

	var p int
	fmt.Scan(&p)
	arr1 := make([]int, p)
	for i := 0; i<p; i++ {
		fmt.Scan(&arr1[i])
	}
	var p2 int
	fmt.Scan(&p2)
	arr2 := make([]int,p2)
	for i := 0; i<p2; i++ {
		fmt.Scan(&arr2[i])
	}

	fmt.Println(solve(n, arr1, arr2))
}
