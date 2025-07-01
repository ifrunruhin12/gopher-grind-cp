/*
https://codeforces.com/contest/1366/problem/A
*/
//Binary search
//greedy
//math

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fast input reader
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func readInt() int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))
	return num
}

func readAB() (int, int) {
	input, _ := reader.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(input))
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	return a, b
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

func main() {
	defer writer.Flush() // Ensure all output is written at once
	n := readInt()

	for i := 0; i < n; i++ {
		a, b := readAB()
		fmt.Fprintln(writer, min((a+b)/3, a, b))
	}
}