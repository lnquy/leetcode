package main

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		temperatures []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"tc1", args{[]int{73, 74, 75, 71, 69, 72, 76, 73}}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"tc2", args{[]int{30, 40, 50, 60}}, []int{1, 1, 1, 0}},
		{"tc3", args{[]int{30, 60, 90}}, []int{1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
