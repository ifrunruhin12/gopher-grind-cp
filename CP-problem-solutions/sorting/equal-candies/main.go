//https://codeforces.com/problemset/problem/1676/B
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
func minCandiesEaten(n int, candies []int) int {
	sum := 0
	lowest := candies[0]
	for i:=0;i<n;i++{
		sum += (candies[i] - lowest)
	}

	return sum
}


func main() {
    defer out.Flush()

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		candies := readInts()
		sort.Ints(candies)
		fmt.Fprintln(out, minCandiesEaten(n, candies))
	}
}
