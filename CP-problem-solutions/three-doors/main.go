//https://codeforces.com/problemset/problem/1709/A
//Implementation
//Greedy

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
    in  = bufio.NewReader(os.Stdin)   // Fast input reader
    out = bufio.NewWriter(os.Stdout)  // Fast output writer
)

func readLine() string {
    line, _ := in.ReadString('\n')               // Reads until newline
    return strings.TrimSpace(line)               // Removes \n and spaces
}

func readInt() int {
    n, _ := strconv.Atoi(readLine())            // Converts line to int
    return n
}

func readInts() []int {
    line := readLine()                           // Reads a line like: "1 2 3"
    parts := strings.Fields(line)                // Splits into ["1", "2", "3"]
    nums := make([]int, len(parts))
    for i, s := range parts {
        nums[i], _ = strconv.Atoi(s)             // Convert each to int
    }
    return nums
}

func main() {
	t := readInt()

	for i:=0;i<t;i++{
		n := readInt()
		nums := readInts()
		
		if nums[n-1] == 0 {
			fmt.Println("NO")
		} else if nums[nums[n-1] - 1] == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
