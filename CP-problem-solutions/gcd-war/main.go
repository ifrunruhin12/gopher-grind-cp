//https://codeforces.com/contest/2104/problem/D

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Node struct {
	val   int64
	left  *Node
	right *Node
}

func main() {
	var l int64 = 1e7
	var isPrime []bool = make([]bool, l)
	for i := int64(2); i < l; i++ {
		isPrime[i] = true
	}
	var pref []int64 = make([]int64, 1)
	for i := int64(2); i < l; i++ {
		if isPrime[i] {
			pref = append(pref, pref[len(pref)-1]+i)
			for j := i * i; j < l; j += i {
				isPrime[j] = false
			}
		}
	}

	in := bufio.NewReader(os.Stdin)
	var tt int
	for fmt.Fscan(in, &tt); tt > 0; tt-- {
		var n int
		fmt.Fscan(in, &n)
		var a []int64 = make([]int64, n)
		sum := int64(0)
		for i := range a {
			fmt.Fscan(in, &a[i])
			sum += a[i]
		}

		ans := 0
		if sum < pref[n] {
			slices.Sort(a)
			for i := range a {
				sum -= a[i]
				if sum >= pref[n-i-1] {
					ans = i + 1
					break
				}
			}
		}

		fmt.Println(ans)
	}
}
