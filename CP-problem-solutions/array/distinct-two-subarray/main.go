//https://www.codechef.com/problems/DIS2SUB
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

//func readStrings() []string {
//    return strings.Fields(readLine())
//}

func allsame(n int, arr []int) bool {
	for i := 1; i < n; i++ {
		if arr[i] != arr[0] {
			return false
		}
	}
	return true
}

func solve(n int, arr []int) int {
	if allsame(n, arr) {
		return -1
	}
	return 2
}

func main() {
    defer out.Flush()

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		arr := readInts()
		fmt.Fprintln(out, solve(n, arr))
	}
}
