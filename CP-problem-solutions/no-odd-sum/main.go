//https://www.codechef.com/problems/NODDSM
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

func solve() {
    N := readInt()
    A := readInts()

    c1 := 0
    for _, v := range A {
        if v == 1 {
            c1++
        }
    }
    c2 := N - c1

    if c1%2 == 1 {
        fmt.Fprintln(out, c2)
    } else {
        costAllOnes := c2
        costAllTwos := c1 / 2
        if costAllOnes < costAllTwos {
            fmt.Fprintln(out, costAllOnes)
        } else {
            fmt.Fprintln(out, costAllTwos)
        }
    }
}

func main() {
    defer out.Flush()
    T := readInt()
    for i := 0; i < T; i++ {
        solve()
    }
}

