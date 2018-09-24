package main

import "testing"

func TestRomanToInteger(t *testing.T) {
	tcs := []struct {
		inRoman string
		outInt  int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},

		{"I", 1},
		{"II", 2},
		{"VII", 7},
		{"V", 5},
		{"X", 10},
		{"XV", 15},
		{"XVIII", 18},
		{"XIX", 19},
		{"XXXIV", 34},
		{"L", 50},
		{"XLIX", 49},
		{"XLVIII", 48},
		{"C", 100},
		{"CCCLXXXIII", 383},
		{"CCXLIV", 244},
		{"D", 500},
		{"CDXXXIII", 433},
		{"DCCCLXXXIII", 883},
		{"M", 1000},
		{"MDCCCLXXXIII", 1883},
		{"MMXVIII", 2018},
		{"MMMCMXCIX", 3999},
	}

	for idx, tc := range tcs {
		out := romanToInt(tc.inRoman)
		if out != tc.outInt {
			t.Errorf("%d. romainToInt(%q) = %d. Expected: %d", idx, tc.inRoman, out, tc.outInt)
		}
	}
}
