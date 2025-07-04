//https://www.codechef.com/problems/YETMON
//sorting

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

//func readStrings() []string {
//    return strings.Fields(readLine())
//}

func main() {
    defer out.Flush()
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		nums := readInts()
		count := 0
		sort.Ints(nums)
		for i := 0; i < n; i++ {
			if nums[i] > count {
				count++
			}
		}
		fmt.Fprintln(out, count)
	}
}
