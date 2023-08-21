package main

import "testing"

func TestIsMonotonic(t *testing.T) {
	tcs := []struct {
		inArr []int
		out   bool
	}{
		{[]int{1, 2, 2, 3}, true},
		{[]int{6, 5, 4, 4}, true},
		{[]int{1, 3, 2}, false},
		{[]int{1, 2, 4, 5}, true},
		{[]int{1, 1, 1}, true},

		{[]int{1}, true},
		{[]int{1, 1}, true},
		{[]int{1, -2}, true},
		{[]int{1, 1, 1, 1, 1, 1, 2}, true},
		{[]int{1, 1, 1, 1, 1, 1, 2, 1}, false},
		{[]int{6, 5, 4, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5, 6}, true},
		{[]int{1, 2, 3, 4, 5, 6, 3}, false},
	}

	for idx, tc := range tcs {
		out := isMonotonic(tc.inArr)
		if out != tc.out {
			t.Errorf("%d. isMonotonic(%v) = %t. Expected: %t", idx, tc.inArr, out, tc.out)
		}
	}
}
