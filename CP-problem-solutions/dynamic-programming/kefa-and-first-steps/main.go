/*
https://codeforces.com/contest/580/problem/A
*/
//brute force
//DP


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
	for i := range nums{
		arr[i], _ = strconv.Atoi(nums[i])
	}

	return arr
}

func longestNonDecreasing(arr []int) int {
    cnt := 1
    max := 1
    for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
            cnt++
        } else {
            if cnt > max {
                max = cnt
            }
            cnt = 1 // Reset cnt to 1 instead of 0 as the current element is still part of a subsegment
        }
    }

    // In case the longest subsegment ends at the last element
    if cnt > max {
        max = cnt
    }

    return max
}

func main() {
	defer writer.Flush()
	arr := readIntArr()
	res := longestNonDecreasing(arr)

	fmt.Fprintln(writer,res)
}