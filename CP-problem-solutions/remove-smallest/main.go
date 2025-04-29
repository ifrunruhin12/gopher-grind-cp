/*
https://codeforces.com/contest/1399/problem/A
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	defer out.Flush()

	t := readInt()
	for i :=0; i < t; i++ {
		n := readInt()
		a := readInts()
		sort.Ints(a)

		check := true
		for j := 1; j<n; j++ {
			if a[j] - a[j-1] > 1 {
				check = false
				break
			}
		}
		if check {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
