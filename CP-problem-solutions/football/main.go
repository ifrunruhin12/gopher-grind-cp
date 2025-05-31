/*
https://codeforces.com/contest/96/problem/A
*/
//Strings

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

func readInts() []int {
	str, _ := reader.ReadString('\n')
	fields := strings.Fields(str)
	nums := make([]int, len(fields))
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}
	return nums
}

func readString() string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}

func stringToList(s string) []string {
	var result []string
	for _, ch := range s {
		result = append(result, string(ch))
	}
	return result
}


func solve(str string) string {
	strarr := stringToList(str)
	cnt0, cnt1 := 0, 0
	for i := 0; i < len(strarr); i++ {
	    if (cnt0 == 7) || (cnt1 == 7) {
			break
		}
		switch {
		case strarr[i] == "0":
			cnt0++
			cnt1 = 0
		case strarr[i] == "1":
			cnt1++
			cnt0 = 0
		}
	}
	if (cnt0 == 7) || (cnt1 == 7) {
		return "YES"
	}
	return "NO"
}

func main() {
	str1 := readString()
	fmt.Println(solve(str1))
}
