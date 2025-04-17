/*
https://codeforces.com/contest/405/problem/A
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

func readInt() int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))
	return num
}

func main() {
	num := readInt()
	sStr, _ := reader.ReadString('\n')
	splitStr := strings.Fields(sStr)

	arr := make([]int, num)
	for i := 0; i<num; i++ {
		arr[i], _ = strconv.Atoi(splitStr[i])
	}

	sort.Ints(arr)
	for j, v := range arr {
		if j > 0{
			fmt.Fprint(writer," ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
