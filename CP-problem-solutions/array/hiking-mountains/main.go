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

func solve(n, k int, arr []int) int {
    totalHikes := 0
    sum := 0
    for i := 0; i < k; i++ {
        sum += arr[i]
    }

    i := 0
    for i <= n-k {
        if sum == 0 {
            totalHikes++
            i += k + 1
            if i <= n-k {
                sum = 0
                for j := i; j < i+k; j++ {
                    sum += arr[j]
                }
            }
        } else {
            if i+k < n {
                sum = sum - arr[i] + arr[i+k]
            }
            i++
        }
    }
    return totalHikes
}




func main() {
    defer out.Flush()

	t := readInt()
	for i := 0; i < t; i++ {
		ints := readInts()
		n, k := ints[0], ints[1]
		arr := readInts()

		fmt.Fprintln(out, solve(n, k, arr))
	}
}
