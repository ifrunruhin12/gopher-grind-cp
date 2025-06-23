//https://codeforces.com/problemset/problem/2114/A

package main

import (
    "bufio"
    "fmt"
    "math"
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

func solve() {
    n := readInt()
    sq := int(math.Ceil(math.Sqrt(float64(n))))
    if sq*sq == n {
        fmt.Fprintf(out, "0 %d\n", sq)
    } else {
        fmt.Fprintln(out, "-1")
    }
}

func main() {
    defer out.Flush()
    t := readInt()
    for i := 0; i < t; i++ {
        solve()
    }
}

