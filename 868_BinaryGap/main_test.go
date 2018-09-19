package main

import "testing"

func TestBinaryGap(t *testing.T) {
	tcs := []struct {
		inNum  int
		outGap int
	}{
		{22, 2},
		{5, 2},
		{6, 1},
		{8, 0},

		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 0},
		{7, 1},
		{9, 3},
		{10, 2},
		{100, 3},
		{128, 0},
		{129, 7},
		{255, 1},
		{256, 0},
		{1024, 0},
		{36565, 4},
		{5000000, 4},
		{1000000000, 3},
	}

	for idx, tc := range tcs {
		out := binaryGap(tc.inNum)
		if out != tc.outGap {
			t.Errorf("%d. binaryGap(%d) = %d. Expected: %d", idx, tc.inNum, out, tc.outGap)
		}
	}
}
