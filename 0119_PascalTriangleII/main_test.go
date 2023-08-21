package main

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tcs := []struct {
		inRowIdx int
		outRow   []int
	}{
		{3, []int{1, 3, 3, 1}},

		{0, []int{1}},
		{1, []int{1, 1}},
		{2, []int{1, 2, 1}},
		{4, []int{1, 4, 6, 4, 1}},
		{5, []int{1, 5, 10, 10, 5, 1}},
		{6, []int{1, 6, 15, 20, 15, 6, 1}},
		{7, []int{1, 7, 21, 35, 35, 21, 7, 1}},
		{8, []int{1, 8, 28, 56, 70, 56, 28, 8, 1}},
		{9, []int{1, 9, 36, 84, 126, 126, 84, 36, 9, 1}},
		{10, []int{1, 10, 45, 120, 210, 252, 210, 120, 45, 10, 1}},
	}

	for idx, tc := range tcs {
		out := getRow(tc.inRowIdx)
		if !reflect.DeepEqual(out, tc.outRow) {
			t.Errorf("%d. getRow(%d) = %v. Expected: %v", idx, tc.inRowIdx, out, tc.outRow)
		}
	}
}
