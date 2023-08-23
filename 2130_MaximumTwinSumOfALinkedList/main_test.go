package main

import "testing"

func Test_pairSum(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"tc1", args{[]int{5, 4, 2, 1}}, 6},
		{"tc2", args{[]int{4, 2, 2, 3}}, 7},
		{"tc3", args{[]int{1, 100000}}, 100001},
		{"tc4", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSum(arrayToLinkedList(tt.args.head)); got != tt.want {
				t.Errorf("pairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
