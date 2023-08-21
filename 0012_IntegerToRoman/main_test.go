package main

import "testing"

func TestName(t *testing.T) {
	tcs := []struct {
		inNum    int
		outRoman string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},

		{1, "I"},
		{2, "II"},
		{5, "V"},
		{7, "VII"},
		{10, "X"},
		{15, "XV"},
		{18, "XVIII"},
		{19, "XIX"},
		{34, "XXXIV"},
		{50, "L"},
		{48, "XLVIII"},
		{49, "XLIX"},
		{100, "C"},
		{244, "CCXLIV"},
		{383, "CCCLXXXIII"},
		{433, "CDXXXIII"},
		{500, "D"},
		{883, "DCCCLXXXIII"},
		{1000, "M"},
		{1883, "MDCCCLXXXIII"},
		{2018, "MMXVIII"},
		{3999, "MMMCMXCIX"},
	}

	for idx, tc := range tcs {
		out := intToRoman(tc.inNum)
		if out != tc.outRoman {
			t.Errorf("%d. intToRoman(%d) = %q. Expected: %q", idx, tc.inNum, out, tc.outRoman)
		}
	}
}
