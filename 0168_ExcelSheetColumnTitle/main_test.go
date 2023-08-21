package main

import (
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	tcs := []struct {
		inNum int
		out   string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},

		{0, ""},
		{2, "B"},
		{26, "Z"},
		{27, "AA"},
		{100, "CV"},
		{500, "SF"},
		{1000, "ALL"},
		{12345, "RFU"},
	}

	for idx, tc := range tcs {
		out := convertToTitle(tc.inNum)
		if out != tc.out {
			t.Errorf("%d. convertToTitle(%d) = %q. Expected: %q", idx, tc.inNum, out, tc.out)
		}
	}
}
