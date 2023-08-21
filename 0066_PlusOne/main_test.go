package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tcs := []struct {
		inArr  []int
		outArr []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},

		{[]int{0}, []int{1}},
		{[]int{5}, []int{6}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{9, 9, 9, 9, 9, 9}, []int{1, 0, 0, 0, 0, 0, 0}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 8}},
	}

	for idx, tc := range tcs {
		out := plusOne(tc.inArr)
		if !reflect.DeepEqual(out, tc.outArr) {
			t.Errorf("%d. plusOne(%v) = %v. Expected: %v", idx, tc.inArr, out, tc.outArr)
		}
	}
}
