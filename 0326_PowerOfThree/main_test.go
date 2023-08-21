package main

import "testing"

var tcs = []struct {
	inNum int
	out   bool
}{
	{27, true},
	{0, false},
	{9, true},
	{45, false},

	{1, true},
	{2, false},
	{3, true},
	{6, false},
	{79, false},
	{81, true},
	{59049, true},
	{1162261467, true},
	{11622614678, false},
	{847288609443, true},
	{8472886094431, false},
	{205891132094649, true},
	{2058911320946491, false},
	{4052555153018976267, true},
	{4052555153018976265, false},
}

func TestIsPowerOfThree(t *testing.T) {
	for idx, tc := range tcs {
		out := isPowerOfThree(tc.inNum)
		if out != tc.out {
			t.Errorf("%d. isPowerOfThree(%d) = %t. Expected: %t", idx, tc.inNum, out, tc.out)
		}
	}
}

func TestIsPowerOfThree_PreCalculated(t *testing.T) {
	for idx, tc := range tcs {
		out := isPowerOfThree_PreCalculated(tc.inNum)
		if out != tc.out {
			t.Errorf("%d. isPowerOfThree_PreCalculated(%d) = %t. Expected: %t", idx, tc.inNum, out, tc.out)
		}
	}
}

var _dump bool

func BenchmarkIsPowerOfThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_dump = isPowerOfThree(i)
	}
}

func BenchmarkIsPowerOfThree_PreCalculated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_dump = isPowerOfThree_PreCalculated(i)
	}
}
