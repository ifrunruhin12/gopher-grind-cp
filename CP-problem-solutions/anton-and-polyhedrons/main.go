/*
https://codeforces.com/contest/785/problem/A
Math
*/


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Define the map (equivalent to Python's dictionary)
	myDict := map[string]int{
		"Tetrahedron":  4,
		"Cube":         6,
		"Octahedron":   8,
		"Dodecahedron": 12,
		"Icosahedron":  20,
	}

	// Read input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text()) // Read number of test cases

	sum := 0
	for i := 0; i < t; i++ {
		scanner.Scan()
		n := scanner.Text()
		sum += myDict[n] // Add the value from the map
	}

	// Print the result
	fmt.Println(sum)
}
