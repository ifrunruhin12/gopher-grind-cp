/*
https://codeforces.com/contest/381/problem/A
*/
//two pointers

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

func solve(n int, arr []int) (int, int) {
	first_pointer := 0
	second_pointer := n - 1
	person1 := 0
	person2 := 0
	cnt := 0
	for cnt < len(arr) {
		if (cnt%2 == 0) && (arr[first_pointer] >= arr[second_pointer]) {
			person1 = person1 + arr[first_pointer]
			first_pointer++
		} else if (cnt%2 == 0) && (arr[first_pointer] < arr[second_pointer]) {
			person1 = person1 + arr[second_pointer]
			second_pointer--
		} else if (cnt%2 == 1) && (arr[first_pointer] >= arr[second_pointer]) {
			person2 = person2 + arr[first_pointer]
			first_pointer++
		} else if (cnt%2 == 1) && (arr[first_pointer] < arr[second_pointer]) {
			person2 = person2 + arr[second_pointer]
			second_pointer--
		}
		cnt++
	}
	return person1, person2
}




func main() {
	defer writer.Flush() // Ensure all output is written at the end
	num := readInt()

	sStr, _ := reader.ReadString('\n')
	splitStr := strings.Fields(sStr)

	s := make([]int, num)
	for i := 0; i < num; i++ {
		s[i], _ = strconv.Atoi(splitStr[i])
	}
	fmt.Println(solve(num,s))
}
