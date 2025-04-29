/*
https://codeforces.com/contest/282/problem/A
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
) 


func main() {
    reader := bufio.NewReader(os.Stdin)
    var n, ans int
    fmt.Fscanln(reader, &n) 

    for i := 0; i < n; i++ {
        s,_ := reader.ReadString('\n')
        s = strings.TrimSpace(s)

        if s == "++X" || s == "X++" {
            ans++
        } else if s == "--X" || s == "X--" {
            ans--
        }
    }

    fmt.Println(ans)
}
