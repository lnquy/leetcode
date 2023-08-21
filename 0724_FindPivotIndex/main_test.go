package main

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"tc1", args{[]int{1, 7, 3, 6, 5, 6}}, 3},
		{"tc2", args{[]int{1, 2, 3}}, -1},
		{"tc3", args{[]int{2, 1, -1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
