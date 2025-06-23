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

func copySlice(s []int) []int {
    cp := make([]int, len(s))
    copy(cp, s)
    return cp
}

func sumArray(arr []int) []int {
    for i := 1; i < len(arr); i++ {
        arr[i] += arr[i-1]
    }
    return arr
}

func main() {
    defer out.Flush()

    n := readInt()
    arr := readInts()

    for i := 0; i < n-1; i++ {
        suffix := copySlice(arr[i:])
        result := sumArray(suffix)
        for j := 0; j < len(result); j++ {
            if j > 0 {
                fmt.Fprint(out, " ")
            }
            fmt.Fprint(out, result[j])
        }
        fmt.Fprintln(out)
    }
}

