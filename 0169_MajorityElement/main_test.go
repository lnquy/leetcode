package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tcs := []struct {
		inArr []int
		out   int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},

		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2}, 1},
		{[]int{1, 2, 1, 2, 2}, 2},
		{[]int{1, 2, 3, 1, 2, 3, 2, 3, 3}, 3},
		{[]int{1, 2, 3, 1, 2, 3, 2, 3, 3, 4, 4, 4, 4, 4}, 4},
		{[]int{1, 2, 3, 1, 2, 4, 4, 4, 4, 4, 3, 2, 3, 3}, 4},
	}

	for idx, tc := range tcs {
		out := majorityElement(tc.inArr)
		if out != tc.out {
			t.Errorf("%d. majorityElement(%v) = %d. Expected: %d", idx, tc.inArr, out, tc.out)
		}
	}
}
