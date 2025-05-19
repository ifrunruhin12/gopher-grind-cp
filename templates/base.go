package main

import (
    "bufio"           // For fast buffered input/output
    "fmt"             // For formatted printing (we only use fmt for testing/debug)
    "os"              // To get access to stdin, stdout
    "strconv"         // To convert strings to ints
    "strings"         // To split strings into slices (fields)
)

//Globals

var (
    in  = bufio.NewReader(os.Stdin)   // Fast input reader
    out = bufio.NewWriter(os.Stdout)  // Fast output writer
)

/*
bufio.NewReader: Buffers stdin so you don’t block on every byte.

bufio.NewWriter: Buffers stdout, much faster. Needs out.Flush() at the end.
*/

func readLine() string {
    line, _ := in.ReadString('\n')               // Reads until newline
    return strings.TrimSpace(line)               // Removes \n and spaces
}

/*
Returns one full line as string.

TrimSpace avoids issues with trailing \n or spaces.
*/

func readInt() int {
    n, _ := strconv.Atoi(readLine())            // Converts line to int
    return n
}

/*
    Calls readLine() → "123"

    Converts to 123

    Ignores errors (_) — fine for CP, but not ideal in prod code.
*/
func readInts() []int {
    line := readLine()                           // Reads a line like: "1 2 3"
    parts := strings.Fields(line)                // Splits into ["1", "2", "3"]
    nums := make([]int, len(parts))
    for i, s := range parts {
        nums[i], _ = strconv.Atoi(s)             // Convert each to int
    }
    return nums
}
//Used for reading multiple ints in one line, space-separated.

func readStrings() []string {
    return strings.Fields(readLine())            // Like readInts, but keeps strings
}

/*
- Splits a line into words by space.

- Perfect for string array input.
*/

func main() {
    defer out.Flush() //Ensures everything buffered in out is printed at the end.


    //code goes here
}

