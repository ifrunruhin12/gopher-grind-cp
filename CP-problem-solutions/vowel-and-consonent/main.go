//https://atcoder.jp/contests/abc049/tasks/abc049_a
//Basic exercise

package main

import "fmt"

func main() {
	var ch byte
	fmt.Scanf("%c", &ch)

	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		fmt.Println("vowel")
	}else {
		fmt.Println("consonant")
	}
}
