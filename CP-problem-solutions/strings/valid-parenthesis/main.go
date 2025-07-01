//https://leetcode.com/problems/valid-parentheses/description/

//stack
//string

package main

import "fmt"

func isValid(s string) bool {
    stack := []rune{}

    for _, char := range s {
        switch char {
        case '(', '{', '[':
            stack = append(stack, char) // push to stack
        case ')', '}', ']':
            if len(stack) == 0 {
                return false // unmatched closing
            }
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1] // pop
            if (char == ')' && top != '(') ||
               (char == '}' && top != '{') ||
               (char == ']' && top != '[') {
                return false // mismatch
            }
        }
    }

    return len(stack) == 0 // if stack is empty, it’s valid
}

func main() {
	test := "[{()}]"
	fmt.Println(isValid(test))
}

