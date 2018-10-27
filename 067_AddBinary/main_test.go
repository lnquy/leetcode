package main

import "testing"

func TestAddBinary(t *testing.T) {
	tcs := []struct {
		inS1   string
		inS2   string
		outStr string
	}{
		{"11", "1", "100"},

		{"0", "0", "0"},
		{"1", "0", "1"},
		{"0", "1", "1"},
		{"1", "1", "10"},
		{"10", "1", "11"},
		{"1", "10", "11"},
		{"10", "10", "100"},
		{"100", "10", "110"},
		{"111", "111", "1110"},
		{"10000000", "10000000", "100000000"},
		{"11111111", "11111111", "111111110"},
	}

	for idx, tc := range tcs {
		out := addBinary(tc.inS1, tc.inS2)
		if out != tc.outStr {
			t.Errorf("%d. addBinary(%q, %q) = %q. Expected: %q", idx, tc.inS1, tc.inS2, out, tc.outStr)
		}
	}
}

var _dump string

func BenchmarkAddBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_dump = addBinary("11111111", "11111111")
	}
}

func BenchmarkAddBinary_abcman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_dump = addBinary_abcman("11111111", "11111111")
	}
}
