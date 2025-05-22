//https://codeforces.com/problemset/problem/1873/B

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


var (
    in  = bufio.NewReader(os.Stdin)   // Fast input reader
    out = bufio.NewWriter(os.Stdout)  // Fast output writer
)


func readLine() string {
    line, _ := in.ReadString('\n')               // Reads until newline
    return strings.TrimSpace(line)               // Removes \n and spaces
}

func readInt() int {
    n, _ := strconv.Atoi(readLine())            // Converts line to int
    return n
}

func readInts() []int {
    line := readLine()                           // Reads a line like: "1 2 3"
    parts := strings.Fields(line)                // Splits into ["1", "2", "3"]
    nums := make([]int, len(parts))
    for i, s := range parts {
        nums[i], _ = strconv.Atoi(s)             // Convert each to int
    }
    return nums
}


//So I have to add 1 to a certain element
//Using bruteforce to add 1 to all elements and check if that's the maximum product

func mul(arr []int) int {
	mul := 1
	for _, v := range arr {
		mul *= v
	}
	return mul
}

//so I have the mul now. and all i need is to add 1 to all element and check the mul and then store the max mul

func main() {
	defer out.Flush()
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		a := readInts()
		maxpro := mul(a)
		for i:=0;i<n;i++{
			a[i]++
			//fmt.Fprintln(out, a)
			maxpro = max(maxpro,mul(a))
			//fmt.Fprintln(out, maxpro)
			a[i]--
			//fmt.Fprintln(out, a)
		}
		fmt.Fprintln(out, maxpro)
	}
}
