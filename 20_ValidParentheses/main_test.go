package main

import "testing"

func TestIsValidParentheses(t *testing.T) {
	tcs := []struct{
		inStr string
		outIsValid bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},

		{"", true},
		{"()", true},
		{"[]", true},
		{"{}", true},
		{"aaa", true},
		{"a(", false},
		{")a(", false},
		{"(abcdr)", true},
		{"([abcdr])", true},
		{"([{abcdr}])", true},
		{"([{abcdr}])", true},
		{"([{}])", true},
		{"asdasd ([{}])", true},
		{"asdasd ([([{}])])", true},
		{"a(a[a(a[a{a}a]a)a]a)a", true},
		{"a(a[a(a[a{a}a]a)a]a)a", true},
		{"{[(])}", false},
		{"{[()}]", false},
		{"a{a[a(a)a}a]a", false},
	}

	for idx, tc := range tcs {
		out := isValid(tc.inStr)
		if out != tc.outIsValid {
			t.Errorf("%d. isValid(%q) = %t. Expected: %t", idx, tc.inStr, out, tc.outIsValid)
		}
	}
}
