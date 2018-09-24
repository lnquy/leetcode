package main

import "testing"

func TestSortArrayByParity(t *testing.T) {
	tcs := []struct {
		inArr []int
	}{
		{[]int{3, 1, 2, 4}},

		{[]int{1}},
		{[]int{1, 2}},
		{[]int{2, 1}},
		{[]int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 10, 11, 13, 15, 17}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 10, 11, 13, 15, 17, 12, 14, 16, 18}},
		{[]int{1, 3, 5, 7, 9}},
		{[]int{2, 4, 6, 8, 10}},
	}

	for idx, tc := range tcs {
		oriInArr := make([]int, len(tc.inArr))
		copy(oriInArr, tc.inArr)
		out := sortArrayByParity(tc.inArr)

		if len(out) != len(oriInArr) {
			t.Errorf("%d. sortArrayByParity(%v) = %v. Expected: array length must be equal", idx, oriInArr, out)
			continue
		}

		if !isParitySorted(out) {
			t.Errorf("%d. sortArrayByParity(%v) = %v. Expected: all even values must be at plcaed the beginning of array",
				idx, oriInArr, out)
		}
	}
}

func isParitySorted(a []int) bool {
	n := len(a)
	if n == 1 {
		return true
	}
	oddFound := false
	for i := 0; i < n; i++ {
		if !oddFound && a[i]%2 == 1 {
			oddFound = true
			continue
		}
		if oddFound && a[i]%2 == 0 {
			return false
		}
	}
	return true
}
