package main

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"tc1", args{[]int{1, 2, 3, 1}}, 2},
		{"tc2", args{[]int{1, 2, 1, 3, 5, 6, 4}}, 5},
		{"tc3", args{[]int{1, 2, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
