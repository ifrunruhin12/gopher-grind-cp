//https://www.codechef.com/practice/course/1to2stars/LP1TO201/problems/IMDB

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

func readInts() []int {
	line := readLine()
	parts := strings.Fields(line)
	nums := make([]int, len(parts))
	for i, s := range parts {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

func main() {
	defer out.Flush()
	t := readInt()
	for i := 0; i < t; i++ {
		arr_num := readInts()
		num1, num2 := arr_num[0], arr_num[1]
		var max_imdb int
		for j := 0; j < num1; j++ {
			arr := readInts()
			if len(arr) >= 2 && arr[0] <= num2 && arr[1] > max_imdb {
				max_imdb = arr[1]
			}
		}
		fmt.Fprintln(out, max_imdb)
	}
}
