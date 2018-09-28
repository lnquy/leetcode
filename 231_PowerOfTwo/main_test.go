package main

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	tcs := []struct {
		inNum int
		out   bool
	}{
		{1, true},
		{16, true},
		{218, false},

		{0, false},
		{-1, false},
		{-8, false},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
		{7, false},
		{8, true},
		{256, true},
		{2000000000, false},
		{2147483648, true},
	}

	for idx, tc := range tcs {
		out := isPowerOfTwo(tc.inNum)
		if out != tc.out {
			t.Errorf("%d. isPowerOfTwo(%d) = %t. Expected: %t", idx, tc.inNum, out, tc.out)
		}
	}

}
