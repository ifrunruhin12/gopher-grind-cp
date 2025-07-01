/*
https://codeforces.com/contest/996/problem/A
*/
//Greedy
//Math

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func readInt() int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))
	return num
}

var arr = [5]int{100, 20, 10, 5, 1} // Reverse order for direct iteration

func main() {
	defer writer.Flush()
	n := readInt()
	cnt := 0

	for _, coin := range arr {
		cnt += n / coin  // Add the quotient directly
		n %= coin        // Get the remaining amount
	}

	fmt.Fprintln(writer, cnt)
}
