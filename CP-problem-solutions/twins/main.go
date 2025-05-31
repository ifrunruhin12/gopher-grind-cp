/*
https://codeforces.com/contest/160/problem/A
*/
//Greedy
//sortings

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func solve(n int, arr []int) int {

	sort.Ints(arr)

	//Reverse the array O(N)
	for i, j := 0, n-1; i < j; i, j = i + 1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	prefixSum := make([]int, n+1)
	for i := 0; i<n; i++ {
		prefixSum[i+1] = prefixSum[i] + arr[i]
	}

	for i := 0; i<n; i++ {
		leftSum := prefixSum[i]
		rightSum := prefixSum[n] - prefixSum[i]
		if leftSum > rightSum {
			return i
		}
	}

	return n

}







func main() {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))


	sStr, _ := reader.ReadString('\n')
	splitStr := strings.Fields(sStr)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i], _ = strconv.Atoi(splitStr[i])

	}

	fmt.Println(solve(n,s))
}
