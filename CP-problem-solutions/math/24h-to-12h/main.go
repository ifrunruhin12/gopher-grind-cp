//https://codeforces.com/problemset/problem/1950/C
//Math

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertTo12Hour(time24 string) string {
	parts := strings.Split(time24, ":")
	hh, _ := strconv.Atoi(parts[0])
	mm := parts[1]
	var period string
	var hour12 int

	if hh == 0 {
		hour12 = 12
		period = "AM"
	} else if hh < 12 {
		hour12 = hh
		period = "AM"
	} else if hh == 12 {
		hour12 = hh
		period = "PM"
	} else {
		hour12 = hh - 12
		period = "PM"
	}

	return fmt.Sprintf("%02d:%s %s", hour12, mm, period)
}

func main() {
	var t int
	fmt.Scan(&t)
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for i := 0; i < t; i++ {
		scanner.Scan()
		time24 := scanner.Text()
		result := convertTo12Hour(strings.TrimSpace(time24))
		fmt.Println(result)
	}
}

