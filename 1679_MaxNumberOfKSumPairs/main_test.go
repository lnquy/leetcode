package main

import "testing"

func Test_maxOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name   string
		args   args
		wantOp int
	}{
		{"tc1", args{[]int{1, 2, 3, 4}, 5}, 2},
		{"tc2", args{[]int{3, 1, 3, 4, 3}, 6}, 1},
		{"tc3", args{[]int{3, 3, 3, 2, 2, 2}, 5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOp := maxOperations(tt.args.nums, tt.args.k); gotOp != tt.wantOp {
				t.Errorf("maxOperations() = %v, want %v", gotOp, tt.wantOp)
			}
		})
	}
}
