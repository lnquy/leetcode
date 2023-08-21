package main

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{"tc1", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"tc2", args{[]int{-1, 1, 0, -3, 3}}, []int{0, 0, 9, 0, 0}},
		{"tc3", args{[]int{1, 2}}, []int{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := productExceptSelf(tt.args.nums); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("productExceptSelf() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
