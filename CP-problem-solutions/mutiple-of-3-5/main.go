//math
//https://projecteuler.net/problem=1
//Math

package main

import "fmt"

func main() {
	sum := 0
	for i:=1;i<1000;i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}
