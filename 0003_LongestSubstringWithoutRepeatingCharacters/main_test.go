package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tcs := []struct {
		inStr     string
		outLength int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"a", 1},
		{"aa", 1},
		{"ab", 2},
		{"aba", 2},
		{"aaaaabcccc", 3},
		{"Given a string, find the length of the longest substring without repeating characters.", 10},
	}

	for idx, tc := range tcs {
		out := lengthOfLongestSubstring(tc.inStr)
		if out != tc.outLength {
			t.Errorf("%d. lengthOfLongestSubstring(%q) = %d. Expected: %d", idx, tc.inStr, out, tc.outLength)
		}
	}
}
