package main

import "testing"

func TestTwoLowerCase(t *testing.T) {
	tcs := []struct {
		inStr  string
		outStr string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},

		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"AbCDefghIJklM", "abcdefghijklm"},
		{"A simple string with UPPERCASE", "a simple string with uppercase"},
		{"Dummy ^#! 1230 text %%   asdT", "dummy ^#! 1230 text %%   asdt"},
	}

	for idx, tc := range tcs {
		out := toLowerCase(tc.inStr)
		if out != tc.outStr {
			t.Errorf("%d. toLowerCase(%q) = %q. Expected: %q", idx, tc.inStr, out, tc.outStr)
		}
	}
}
