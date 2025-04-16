/*
https://codeforces.com/contest/58/problem/A
*/

package main

import "fmt"

func main() {
    var s string
    fmt.Scanln(&s)

    target := "hello"
    targetIndex := 0

    // Iterate over the string s
    for i := 0; i < len(s); i++ {
        if s[i] == target[targetIndex] {
            targetIndex++
        }
        // If we have matched all characters of "hello"
        if targetIndex == len(target) {
            break
        }
    }

    if targetIndex == len(target) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
