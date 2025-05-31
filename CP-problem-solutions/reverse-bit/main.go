//https://leetcode.com/problems/reverse-bits/description/
//

package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= num & 1
		num >>= 1
	}
	return result
}

func main() {
	fmt.Println(reverseBits(43261596))     // Output: 964176192
	fmt.Println(reverseBits(4294967293))   // Output: 3221225471
}

