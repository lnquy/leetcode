package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad") == "bab" || longestPalindrome("babad") == "aba")
	fmt.Println(longestPalindrome("cbbd") == "bb")
}

// O(n^2)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		exp1, exp2 := expandOnChar(s, i, i), expandOnChar(s, i, i+1)
		if m := max(exp1, exp2); m > (end - start) {
			start = i - (m-1)/2
			end = i + m/2
		}
	}
	return s[start : end+1]
}

// O(n)
func expandOnChar(s string, start, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return end - start - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
