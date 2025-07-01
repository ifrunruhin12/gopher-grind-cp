//https://leetcode.com/problems/binary-search/description/
//binary search

package main

import "fmt"

func solution(slc []int, target int) int {
	up := len(slc) - 1 
	down := 0 
	for down <= up {
		index := (up + down) / 2 
		if slc[index] == target {
			return index
		}else if slc[index] < target {
			down = index + 1
		}else {
			up = index - 1
		}
	}
	return -1	
}


func main(){
	s := []int{-1, 0, 2, 4, 6, 8}
	a := 4
	res := solution(s, a)
	fmt.Println(res)
}
