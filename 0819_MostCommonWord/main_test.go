package main

import "testing"

func TestMostCommonWord(t *testing.T) {
	tcs := []struct {
		inParagraph string
		inBanned    []string
		out         string
	}{
		{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"},

		{"a", []string{""}, "a"},
		{"a b a", []string{""}, "a"},
		{"a b a", []string{"a"}, "b"},
		{"c  a  b    a c   d", []string{"a"}, "c"},
		{"!c;  a.  b,    a' !ca?c   d;", []string{"a"}, "c"},
	}

	for idx, tc := range tcs {
		out := mostCommonWord(tc.inParagraph, tc.inBanned)
		if out != tc.out {
			t.Errorf("%d. mostCommonWord(%q, %v) = %q. Expected: %q", idx, tc.inParagraph, tc.inBanned, out, tc.out)
		}
	}
}
