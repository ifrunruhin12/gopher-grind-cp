//https://codeforces.com/contest/2109/problem/A
//Implementation

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

func main() {
	defer out.Flush()
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		arr := readInts()

		allOnes := true
		hasConsecutiveZeros := false

		for j := 0; j < n; j++ {
			if arr[j] == 0 {
				allOnes = false
			}
			if j > 0 && arr[j] == 0 && arr[j-1] == 0 {
				hasConsecutiveZeros = true
			}
		}

		if allOnes || hasConsecutiveZeros {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
