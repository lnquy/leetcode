package main

import "testing"

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name        string
		args        args
		wantLongest int
	}{
		{"tc1", args{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2}, 6},
		{"tc2", args{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLongest := longestOnes(tt.args.nums, tt.args.k); gotLongest != tt.wantLongest {
				t.Errorf("longestOnes() = %v, want %v", gotLongest, tt.wantLongest)
			}
		})
	}
}
