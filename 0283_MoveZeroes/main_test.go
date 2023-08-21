package main

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"tc1", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
		{"tc2", args{[]int{0}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes(), got %v, expected %v", tt.args.nums, tt.want)
			}
		})
	}
}
