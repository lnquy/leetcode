package main

import "testing"

func TestHasAlternatingBits(t *testing.T) {
	tcs := []struct {
		inNum int
		out   bool
	}{
		{5, true},
		{7, false},
		{11, false},
		{10, true},

		{0, true},
		{1, true},
		{2, true},
		{3, false},
		{4, false},
		{6, false},
		{8, false},
		{9, false},
		{21, true},
		{42, true},
		{85, true},
		{170, true},
		{43690, true},
		{174762, true},
		{11184810, true},
		{2863311530, true},
	}

	for idx, tc := range tcs {
		out := hasAlternatingBits(tc.inNum)
		if out != tc.out {
			t.Errorf("%d. hasAlternatingBits(%d) = %t. Expected: %t", idx, tc.inNum, out, tc.out)
		}
	}
}

var _dump bool // Use global var to prevent compiler optimizes dead code

func BenchmarkHasAlternativeBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_dump = hasAlternatingBits(i)
	}
}
