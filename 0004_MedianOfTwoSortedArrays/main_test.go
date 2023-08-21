package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tcs := []struct {
		inNums1   []int
		inNums2   []int
		outMedian float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},

		{[]int{}, []int{1}, 1},
		{[]int{1}, []int{}, 1},
		{[]int{1}, []int{1}, 1},
		{[]int{1}, []int{2}, 1.5},
		{[]int{1, 2}, []int{}, 1.5},
		{[]int{}, []int{1, 2}, 1.5},
		{[]int{1}, []int{1, 2}, 1},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9}, 5},
		{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}, 5.5},
		{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10, 12}, 6},
	}

	for idx, tc := range tcs {
		out := findMedianSortedArrays_Slow(tc.inNums1, tc.inNums2)
		if out != tc.outMedian {
			t.Errorf("%d. findMedianSortedArrays(%v, %v) = %f. Expected: %f", idx, tc.inNums1, tc.inNums2, out, tc.outMedian)
		}
	}
}
