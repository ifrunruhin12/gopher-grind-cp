/*
https://codeforces.com/contest/1690/problem/D
*/
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

func readString() string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}

func minWhiteBlocks(n, k int, s string) int {
	runes := []rune(s) // Use []rune instead of []string for better efficiency

	// Count "W" in the first k-length window
	currW := 0
	for i := 0; i < k; i++ {
		if runes[i] == 'W' {
			currW++
		}
	}

	// Sliding window approach
	minW := currW
	for i := k; i < n; i++ {
		if runes[i] == 'W' { // Add new character entering the window
			currW++
		}
		if runes[i-k] == 'W' { // Remove old character leaving the window
			currW--
		}
		if currW < minW {
			minW = currW
		}
	}

	return minW
}

func main() {
	defer writer.Flush() // Ensure output is written at once

	t := readInt()
	for i := 0; i < t; i++ {
		n, k := readAB()
		str := readString()
		fmt.Fprintln(writer, minWhiteBlocks(n, k, str))
	}
}