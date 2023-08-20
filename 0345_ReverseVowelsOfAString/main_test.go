package main

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"tc1", args{"hello"}, "holle"},
		{"tc2", args{"leetcode"}, "leotcede"},
		{"tc3", args{"ai"}, "ia"},
		{"tc4", args{" apG0i4maAs::sA0m4i0Gp0"}, " ipG0A4mAas::si0m4a0Gp0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
