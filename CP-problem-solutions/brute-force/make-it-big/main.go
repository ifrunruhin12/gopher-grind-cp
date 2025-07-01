//https://toph.co/p/make-it-big

package main

import (
	"fmt"
	"strconv"
)

func swap(d []rune, i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var N int
		fmt.Scan(&N)

		maxNum := N

		digits := []rune(strconv.Itoa(N))
		n := len(digits)

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				swap(digits, i, j)
				for k := 0; k < n; k++ {
					for l := k + 1; l < n; l++ {
						swap(digits, k, l)
						num, _ := strconv.Atoi(string(digits))
						if num > maxNum {
							maxNum = num
						}
						swap(digits, k, l)
					}
				}
				swap(digits, i, j)
			}
		}

		fmt.Println(maxNum)
	}
}
