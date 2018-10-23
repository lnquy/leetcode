package main

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	tcs := []struct {
		inImg  [][]int
		outImg [][]int
	}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},

		{[][]int{}, [][]int{}},
		{[][]int{{1}}, [][]int{{0}}},
		{[][]int{{1, 0}}, [][]int{{1, 0}}},
		{[][]int{{0, 1}}, [][]int{{0, 1}}},
		{[][]int{{1, 0, 1}}, [][]int{{0, 1, 0}}},
		{[][]int{{0, 1, 0}}, [][]int{{1, 0, 1}}},
		{[][]int{{0, 0, 0}}, [][]int{{1, 1, 1}}},
		{[][]int{{1, 1, 1}}, [][]int{{0, 0, 0}}},
	}

	for idx, tc := range tcs {
		tmp := make([][]int, len(tc.inImg))
		copy(tmp, tc.inImg)
		out := flipAndInvertImage(tc.inImg)
		if !reflect.DeepEqual(out, tc.outImg) {
			t.Errorf("%d. flipAndInvertImage(%v) = %v. Expected: %v", idx, tmp, out, tc.outImg)
		}
	}
}
