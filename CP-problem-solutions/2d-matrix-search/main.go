package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

    m := len(matrix)
    n := len(matrix[0])
    left, right := 0, m*n-1

    for left <= right {
        mid := (left + right) / 2
        row := mid / n
        col := mid % n
        midVal := matrix[row][col]

        if midVal == target {
            return true
        } else if midVal < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return false
}

func main() {
	s := [][]int{{1,2},{3,4}}
	p := 2
	res := searchMatrix(s,p)
	fmt.Println(res)
}
