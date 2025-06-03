//https://www.codechef.com/problems/SUBSUM3
//Array
//Brute force

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

//func readStrings() []string {
//    return strings.Fields(readLine())
//}

func main() {
    defer out.Flush()

	t := readInt()
	for ; t > 0; t-- {
		n := readInt()
		arr := readInts()
		cnt := 0

		for i := 0; i < n; i++ {
			sum := 0
			for j := i; j < n; j++ {
				sum += arr[j]
				if sum % 3 == 0 {
					cnt++
				}			
			}
		}

		if cnt > 0 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
