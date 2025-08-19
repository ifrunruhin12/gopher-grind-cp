//https://codeforces.com/contest/2131/problem/A

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
	line := readLine()
	parts := strings.Fields(line)
	nums := make([]int, len(parts))
	for i, s := range parts {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

func solve(arr1, arr2 []int) int {
	sum := 0
	for i := range arr1 {
		sum += max(arr1[i]-arr2[i], 0)
	}
	return sum + 1
}

func main() {
	defer out.Flush()
	t := readInt()
	for range t {
		n := readInt()
		_ = n
		arr1 := readInts()
		arr2 := readInts()
		fmt.Fprintln(out, solve(arr1, arr2))
	}
}
