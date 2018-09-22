package main

import "testing"

func TestRemoveElement(t *testing.T) {
	tcs := []struct {
		inArr  []int
		inVal  int
		outLen int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},

		{[]int{}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{1}, 0, 1},
		{[]int{1, 2, 3}, 0, 3},
		{[]int{1, 1, 2, 3}, 1, 2},
		{[]int{1, 2, 2, 3}, 2, 2},
		{[]int{1, 2, 3, 3}, 3, 2},
		{[]int{5, 0, 5, 1, 5, 2, 5, 3, 5, 4, 5, 5, 6, 5, 7, 5, 8, 5, 9, 5}, 5, 9},
	}

	for idx, tc := range tcs {
		oriInArr := make([]int, len(tc.inArr))
		copy(oriInArr, tc.inArr)

		out := removeElement(tc.inArr, tc.inVal)
		if out != tc.outLen {
			t.Errorf("%d. removeElement(%v, %d) = %d. Expected: %d", idx, oriInArr, tc.inVal, out, tc.outLen)
			continue
		}
		for i := 0; i < out; i++ {
			if tc.inArr[i] == tc.inVal {
				t.Errorf("%d. removeElement(%v, %d) = %v. Contains element(s) which must be removed", idx, oriInArr, tc.inVal, tc.inArr)
				break
			}
		}
	}
}
