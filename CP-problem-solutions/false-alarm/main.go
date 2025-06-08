//https://codeforces.com/contest/2117/problem/A
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

func solve(n, x int, arr []int) string {
	first_door := 0
	last_door := n - 1
	for i := 0; i < n; i++ {
		if arr[i] == 1 {
			first_door = i
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		if arr[i] == 1 {
			last_door = i
			break
		}
	}
	diff := (last_door - first_door) + 1
	if diff <= x {
		return "YES"
	}
	return "NO"
}

func main() {
    defer out.Flush()
    t := readInt()
	for i := 0; i < t; i++ {
		ints := readInts()
		n, x := ints[0], ints[1]
		arr := readInts()
		fmt.Println(solve(n, x, arr))
	}
}
