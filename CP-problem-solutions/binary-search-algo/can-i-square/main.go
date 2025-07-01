//https://codeforces.com/problemset/problem/1915/C

//Binary search

package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
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


func isPerfectSquare(n int) bool {
    if n < 0 {
        return false
    }
    root := int(math.Sqrt(float64(n)))
    return root*root == n
}

func main() {
	defer out.Flush()
	t := readInt()
	for i:=0;i<t;i++{
		n := readInt()
		arr := readInts()
		sum := 0
		for j:=0;j<n;j++{
			sum = sum + arr[j]
		}
		if isPerfectSquare(sum) {
			fmt.Fprintln(out, "YES")
		}else {
			fmt.Fprintln(out, "No")
		}
	}
}
