package main

import "testing"

func TestHammingWeight(t *testing.T) {
	tcs := []struct {
		inNum   uint32
		outBits int
	}{
		{11, 3},
		{128, 1},

		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{6, 2},
		{255, 8},
		{524287, 19},
		{524287, 19},
		{87381, 9},
	}

	for idx, tc := range tcs {
		out := hammingWeight(tc.inNum)
		if out != tc.outBits {
			t.Errorf("%d. hammingWeight(%d) = %d. Expected: %d", idx, tc.inNum, out, tc.outBits)
		}
	}
}
