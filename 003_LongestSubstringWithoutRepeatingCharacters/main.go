package main

import "fmt"

func main() {
	fmt.Println("abcabcbb", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbbb", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstring("pwwkew"))
}

// O(n)
func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)

	lastDupIdx, length := -1, 0
	for i, v := range s {
		if idx, ok := m[string(v)]; ok { // Duplicated character
			lastDupIdx = max(lastDupIdx, idx)
		}

		length = max(length, i - lastDupIdx)
		m[string(v)] = i
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
