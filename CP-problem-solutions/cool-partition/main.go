//https://codeforces.com/contest/2117/problem/C

//Prefix Sum

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	line, _ := in.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInt() int {
	n, _ := strconv.Atoi(readLine())
	return n
}

func readInts() []int {
	parts := strings.Fields(readLine())
	nums := make([]int, len(parts))
	for i, p := range parts {
		nums[i], _ = strconv.Atoi(p)
	}
	return nums
}

func solve(n int, a []int) int {
	seen := make(map[int]bool)
	totalUnique := make(map[int]bool)
	count := 1

	for i := 0; i < n; i++ {
		val := a[i]
		seen[val] = true
		totalUnique[val] = true

		if len(seen) == len(totalUnique) {
			count++
			seen = make(map[int]bool)
		}
	}

	return count - 1
}

func main() {
	defer out.Flush()
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		a := readInts()
		fmt.Fprintln(out, solve(n, a))
	}
}
