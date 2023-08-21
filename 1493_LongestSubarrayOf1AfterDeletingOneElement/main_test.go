package main

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name        string
		args        args
		wantLongest int
	}{
		{"tc1", args{[]int{1, 1, 0, 1}}, 3},
		{"tc2", args{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}}, 5},
		{"tc3", args{[]int{1, 1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLongest := longestSubarray(tt.args.nums); gotLongest != tt.wantLongest {
				t.Errorf("longestSubarray() = %v, want %v", gotLongest, tt.wantLongest)
			}
		})
	}
}
