package main

import "testing"

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name       string
		args       args
		wantMaxAvg float64
	}{
		{"tc1", args{[]int{1, 12, -5, -6, 50, 3}, 4}, 12.75},
		{"tc2", args{[]int{5}, 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMaxAvg := findMaxAverage(tt.args.nums, tt.args.k); gotMaxAvg != tt.wantMaxAvg {
				t.Errorf("findMaxAverage() = %v, want %v", gotMaxAvg, tt.wantMaxAvg)
			}
		})
	}
}
