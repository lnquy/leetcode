package main

import "testing"

func TestCountSegments(t *testing.T) {
	tcs := []struct {
		inStr       string
		outSegments int
	}{
		{"Hello, my name is John", 5},

		{"", 0},
		{" ", 0},
		{"    ", 0},
		{"a", 1},
		{"  a  ", 1},
		{"a b", 2},
		{"a     b", 2},
		{"    a     b   ", 2},
		{"    a     b   c   ", 3},
		{"    awqerq,     asdqb;   cqweqf   asdaswq", 4},
	}

	for idx, tc := range tcs {
		out := countSegments(tc.inStr)
		if out != tc.outSegments {
			t.Errorf("%d. countSegments(%q) = %d. Expected: %d", idx, tc.inStr, out, tc.outSegments)
		}
	}
}
