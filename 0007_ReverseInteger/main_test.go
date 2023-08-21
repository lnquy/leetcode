package main

import "testing"

func TestReverse(t *testing.T) {
	tcs := []struct {
		inInt  int
		outInt int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},

		{0, 0},
		{1, 1},
		{-1, -1},
		{1000, 1},
		{-1000, -1},
		{12345, 54321},
		{2147483647, 0},
		{-2147483648, 0},
		{2147483641, 1463847412},
	}

	for idx, tc := range tcs {
		out := reverse(tc.inInt)
		if out != tc.outInt {
			t.Errorf("%d. reverse(%d) = %d. Expected: %d", idx, tc.inInt, out, tc.outInt)
		}
	}
}
