//https://codeforces.com/contest/2117/problem/B
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

func solve(n int) []int {
    res := make([]int, 0, n)
    for i := 2; i <= n; i += 2 {
        res = append(res, i)
    }
    start := n
    if start%2 == 0 {
        start--
    }
    for i := start; i >= 1; i -= 2 {
        res = append(res, i)
    }

    return res
}

func main() {
    defer out.Flush()
    t := readInt()
    for i := 0; i < t; i++ {
        n := readInt()
        res := solve(n)
        for j := 0; j < len(res); j++ {
            fmt.Fprintf(out, "%d ", res[j])
        }
        fmt.Fprintln(out)
    }
}
