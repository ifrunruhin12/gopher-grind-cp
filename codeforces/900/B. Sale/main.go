/*
https://codeforces.com/contest/34/problem/B
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Fast input reader
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func readInts() []int {
	str, _ := reader.ReadString('\n')
	fields := strings.Fields(str)
	nums := make([]int, len(fields))
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}
	return nums
}


func readAB() (int, int) {
	input, _ := reader.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(input))
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	return a, b
}

func driver(num int) int {
	if num < 0 {
		return -num
	}
	return 0
}

func main() {
	defer writer.Flush() // Ensure all output is written at the end
	a, b := readAB()

	sStr, _ := reader.ReadString('\n')
	splitStr := strings.Fields(sStr)
	s := make([]int, a)
	for i := 0; i < a; i++ {
		s[i], _ = strconv.Atoi(splitStr[i])
	}
	sort.Ints(s)
	cnt := 0
	for i := 0; i < b; i++ {
		cnt = cnt + driver(s[i])
	}
	fmt.Fprintln(writer,cnt)

}