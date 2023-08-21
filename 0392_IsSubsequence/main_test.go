package main

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"tc1", args{"abc", "ahbgdc"}, true},
		{"tc2", args{"axc", "ahbgdc"}, false},
		{"tc3", args{"", "ahbgdc"}, true},
		{"tc4", args{"asd", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
