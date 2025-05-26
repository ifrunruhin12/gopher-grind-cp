func longestPalindrome(s string) string {
    if len(s) < 1 {
        return ""
    }

    start, end := 0, 0

    for i := 0; i < len(s); i++ {
        len1 := expandAroundCenter(s, i, i)     // Odd length palindrome
        len2 := expandAroundCenter(s, i, i+1)   // Even length palindrome
        maxLen := max(len1, len2)

        if maxLen > end-start {
            start = i - (maxLen-1)/2
            end = i + maxLen/2
        }
    }

    return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1 // actual length of the palindrome
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

