package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tcs := []struct {
		inStr  string
		outLen int
	}{
		{"Hello World", 5},

		{"abc   wqreqwf", 7},
		{"abcdef", 6},
		{"abcd ", 4},
		{"abcd   ", 4},
	}

	for idx, tc := range tcs {
		out := lengthOfLastWord(tc.inStr)
		if out != tc.outLen {
			t.Errorf("%d. lengthOfLastWord(%q) = %d. Expected: %d", idx, tc.inStr, out, tc.outLen)
		}
	}
}
