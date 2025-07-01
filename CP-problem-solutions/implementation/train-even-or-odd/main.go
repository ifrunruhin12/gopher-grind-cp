//https://www.codechef.com/problems/TRAINEVOD

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

func readStrings() []string {
    return strings.Fields(readLine())
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
    defer out.Flush()

	t := readInt()

	for i := 0; i < t; i++ {
		n := readInt()
		arr := readInts()
		even, odd := 0, 0

		for j:=0; j < n; j++ {
			if j % 2 == 0 {
				odd += arr[j]
			} else {
				even += arr[j]
			}
		}

		res := maxNum(even, odd)
		fmt.Fprintln(out, res)
	}
}
