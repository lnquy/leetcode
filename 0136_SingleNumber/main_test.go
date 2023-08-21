package main

import "testing"

func TestSingleNumber(t *testing.T) {
	tcs := []struct {
		inNums []int
		outNum int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},

		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 2, 1}, 2},
		{[]int{2, 1, 1}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 6, 5, 4, 3, 2}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8}, 9},
	}

	for idx, tc := range tcs {
		out := singleNumber(tc.inNums)
		if out != tc.outNum {
			t.Errorf("%d. singleNumber(%v) = %d. Expected: %d", idx, tc.inNums, out, tc.outNum)
		}
	}
}
