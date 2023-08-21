package main

import "testing"

func TestTwoSum(t *testing.T) {
	tcs := []struct {
		inNums   []int
		inTarget int
		outIdxs  []int
	}{
		{[]int{10, 20}, 30, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
		{[]int{2, 7, 11, 15}, 18, []int{1, 2}},
		{[]int{2, 7, 11, 15}, 26, []int{2, 3}},
	}

	for idx, tc := range tcs {
		out := twoSum(tc.inNums, tc.inTarget)
		if (out[0] == tc.outIdxs[0] && out[1] == tc.outIdxs[1]) || (out[0] == tc.outIdxs[1] && out[1] == tc.outIdxs[0]) {
			continue
		}
		t.Errorf("%d. twoSum(%v, %d) = %v. Expected: %v", idx, tc.inNums, tc.inTarget, out, tc.outIdxs)
	}
}
