package main

import "testing"

func TestUniqueMorseRepresentations(t *testing.T) {
	tcs := []struct {
		inStrs []string
		out    int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},

		{[]string{}, 0},
		{[]string{""}, 1},
		{[]string{"a"}, 1},
		{[]string{"a", "a", "a"}, 1},
		{[]string{"a", "b", "a"}, 2},
		{[]string{"a", "b", "c"}, 3},
	}

	for idx, tc := range tcs {
		out := uniqueMorseRepresentations(tc.inStrs)
		if out != tc.out {
			t.Errorf("%d. uniqueMorseRepresentations(%q) = %d. Expected: %d", idx, tc.inStrs, out, tc.out)
		}
	}
}
