package main

import "testing"

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{"tc1", args{"abciiidef", 3}, 3},
		{"tc2", args{"aeiou", 2}, 2},
		{"tc3", args{"leetcode", 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxVowels(tt.args.s, tt.args.k); gotMax != tt.wantMax {
				t.Errorf("maxVowels() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
