package main

import "testing"

func TestSearch(t *testing.T) {
	tcs := []struct {
		inNums   []int
		inTarget int
		outIdx   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},

		{[]int{}, 2, -1},
		{[]int{2}, 2, 0},
		{[]int{1, 2}, 2, 1},
		{[]int{1, 3, 5, 7, 9, 11}, 1, 0},
		{[]int{1, 3, 5, 7, 9, 11}, 3, 1},
		{[]int{1, 3, 5, 7, 9, 11}, 5, 2},
		{[]int{1, 3, 5, 7, 9, 11}, 7, 3},
		{[]int{1, 3, 5, 7, 9, 11}, 9, 4},
		{[]int{1, 3, 5, 7, 9, 11}, 11, 5},
		{[]int{1, 3, 5, 7, 9, 11}, 2, -1},
		{[]int{1, 3, 5, 7, 9, 11}, 8, -1},
	}

	for idx, tc := range tcs {
		out := search(tc.inNums, tc.inTarget)
		if out != tc.outIdx {
			t.Errorf("%d. search(%v, %d) = %d. Expected: %d", idx, tc.inNums, tc.inTarget, out, tc.outIdx)
		}
	}
}
