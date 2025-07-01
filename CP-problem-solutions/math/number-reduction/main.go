//https://www.codechef.com/START190D/problems/RED23

package main

import "fmt"

func main() {
	var t, n int
	fmt.Scan(&t)
	for i:=0;i<t;i++{
		fmt.Scan(&n)
		for n > 1 {
			if n%2 == 0 {
				n = n/2
			}else if n%2 == 1 && n > 3 {
				n = n - 3
			}else {
				break
			}
		}
		fmt.Println(n)
	}
}

