//https://codeforces.com/contest/2125/problem/A

package main

import (
	"bufio"
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

func main() {
	defer out.Flush()

	t := readInt()

	for range t {
		s := readLine()

		cnt := 0
		sb := strings.Builder{}
		sb.Grow(len(s))

		for _, ch := range s {
			if ch == 'T' {
				cnt++
			} else {
				sb.WriteRune(ch)
			}
		}

		result := strings.Repeat("T", cnt) + sb.String()
		out.WriteString(result + "\n")
	}
}
