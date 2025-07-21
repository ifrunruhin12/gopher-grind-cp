//https://codeforces.com/contest/2126/problem/A
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

func smallestDigit(n int) int {
	smallest := 9
	for n > 0 {
		digit := n % 10
		if digit < smallest {
			smallest = digit
		}
		n /= 10
	}
	return smallest
}


func main() {
    defer out.Flush()
    
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		fmt.Fprintln(out, smallestDigit(n))
	}
}
