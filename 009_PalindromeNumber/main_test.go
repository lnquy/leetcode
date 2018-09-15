package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tcs := []struct{
		inNum int
		outPalindrome bool
	}{
		{121, true},
		{-121, false},
		{10, false},

		{0, true},
		{1, true},
		{50000, false},
		{55, true},
		{33333, true},
		{1234, false},
		{1234567, false},
		{123321, true},
		{123454321, true},
		{3000000000, false},
		{3000000003, true},
	}

	for idx, tc := range tcs {
		out := isPalindrome(tc.inNum)
		if out != tc.outPalindrome {
			t.Errorf("%d. isPalindrome(%d) = %t. Expected: %t", idx, tc.inNum, out, tc.outPalindrome)
		}
	}
}
