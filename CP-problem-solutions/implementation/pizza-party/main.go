//https://www.codechef.com/problems/PIZZAPARTY

package main

import "fmt"

func main() {
	var b, g int
	fmt.Scan(&b, &g)
	b = b + 1
	pizza := (b * 4) + (g * 3)
	flt_value := float64(pizza) / 8
	int_value := pizza / 8
	if flt_value == float64(int_value) {
		fmt.Println(int_value)
	} else {
		fmt.Println(int_value + 1)
	}
}
