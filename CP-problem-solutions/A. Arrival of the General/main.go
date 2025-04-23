/*
https://codeforces.com/contest/144/problem/A
*/

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

//Function to read an array of n integers
func readIntArr() []int {
	// Read the first integer (size of the array)
	sizeStr, _ := reader.ReadString('\n')
	size, _ := strconv.Atoi(strings.TrimSpace(sizeStr))

	// Read the space-separated integers
	numStrs, _ := reader.ReadString('\n')
	nums := strings.Fields(numStrs)

	// Convert to integer array
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(nums[i])
	}

	return arr
}

func minOfList(arr []int) int {
	if len(arr) == 0 {
		return -1 // Return -1 if the list is empty
	}

	minIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func maxOfList(arr []int) int {
	if len(arr) == 0 {
		return -1 // Return -1 if the list is empty
	}

	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}


func main() {
	defer writer.Flush()
	var res int

	arr := readIntArr()
	l := len(arr) - 1
	minA := minOfList(arr)
	maxA := maxOfList(arr)

	if minA == maxA {
		res = 0
	} else if maxA < minA {
		res = maxA + (l - minA)
	} else {
		res = maxA + (l - minA) - 1
	}

	fmt.Fprintln(writer,res)
}
