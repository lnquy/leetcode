package main

import "testing"

func TestDistributeCandies(t *testing.T) {
	tcs := []struct {
		inCandies []int
		out       int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},

		{[]int{1, 1}, 1},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{1, 2, 3, 4, 1, 1}, 3},
		{[]int{1, 2, 3, 4, 1, 1, 1, 1}, 4},
		{[]int{1, 2, 3, 4, 1, 1, 1, 1, 1, 1}, 4},
		{[]int{1, 2, 3, 4, 1, 1, 1, 1, 1, 1, 5, 5}, 5},
		{[]int{1, 2, 3, 4, 1, 1, 1, 1, 1, 1, 5, 6}, 6},
	}

	for idx, tc := range tcs {
		out := distributeCandies(tc.inCandies)
		if out != tc.out {
			t.Errorf("%d. distributeCandies(%v) = %d. Expected: %d", idx, tc.inCandies, out, tc.out)
		}
	}
}
