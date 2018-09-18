package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tcs := []struct {
		inArr  []int
		outDup int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},

		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 1},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5}, 5},
		{[]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 6, 6, 7, 7, 8}, 8},
	}

	for idx, tc := range tcs {
		out := removeDuplicates(tc.inArr)
		if out != tc.outDup {
			t.Errorf("%d. removeDuplicates(%v) = %d. Expected: %d", idx, tc.inArr, out, tc.outDup)
		}
	}
}
