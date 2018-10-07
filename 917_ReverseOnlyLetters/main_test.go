package main

import "testing"

func TestReverseLetterOnly(t *testing.T) {
	tcs := []struct {
		inStr string
		out   string
	}{
		{"ab-cd", "dc-ba"},
		{"a-bC-dEf-ghIj", "j-Ih-gfE-dCba"},
		{"Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"},

		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"a bc", "c ba"},
		{"ab!c d ~ 0 efgh I J-k ", "kJ!I h ~ 0 gfed c b-a "},
	}

	for idx, tc := range tcs {
		out := reverseOnlyLetters(tc.inStr)
		if out != tc.out {
			t.Errorf("%d. reverseLetterOnly(%q) = %q. Expected: %q", idx, tc.inStr, out, tc.out)
		}
	}
}
