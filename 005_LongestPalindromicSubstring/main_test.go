package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tcs := []struct{
		inStr string
		outStr string
	}{
		{"babad", "aba"},
		{"cbbd", "bb"},

		{"", ""},
		{"a", "a"},
		{"aa", "aa"},
		{"aaa", "aaa"},
		{"aaba", "aba"},
		{"aaba", "aba"},
		{"abba", "abba"},
		{"abba", "abba"},
		{"cbbabb", "bbabb"},
		{"abcdef", "a"},
		{"abcdefedcba", "abcdefedcba"},
	}

	for idx, tc := range tcs {
		out := longestPalindrome(tc.inStr)
		if len(tc.outStr) != len(out) {
			t.Errorf("%d. longestPalindrome(%q) = %q. Expected: %q", idx, tc.inStr, out, tc.outStr)
			continue
		}
		if out != tc.outStr {
			t.Logf("[WARN] Longest length of palidrome is true but returned string not matched:\n   >> %d. longestPalindrome(%q) = %q. Expected: %q", idx, tc.inStr, out, tc.outStr)
		}

	}
}
