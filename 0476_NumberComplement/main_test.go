package main

import "testing"

func TestFindComplement(t *testing.T) {
	tcs := []struct {
		inNum  int
		outNum int
	}{
		{5, 2},
		{1, 0},

		{0, 1},
		{2, 1},
		{3, 0},
		{4, 3},
		{6, 1},
		{7, 0},
		{8, 7},
		{9, 6},
		{10, 5},
		{128, 127},
		{255, 0},
		{65535, 0},
		{65536, 65535},
		{16777215, 0},
		{16777216, 16777215},
		{33554431, 0},
		{33554432, 33554431},
		{4294967295, 0},
		{682, 341},
		{89478485, 44739242.},
	}

	for idx, tc := range tcs {
		out := findComplement(tc.inNum)
		if out != tc.outNum {
			t.Errorf("%d. findComplement(%d) = %d. Expected: %d", idx, tc.inNum, out, tc.outNum)
		}
	}
}
