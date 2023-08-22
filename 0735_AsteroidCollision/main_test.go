package main

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"tc1", args{[]int{5, 10, -5}}, []int{5, 10}},
		{"tc2", args{[]int{8, -8}}, []int{}},
		{"tc3", args{[]int{10, 2, -5}}, []int{10}},
		{"tc4", args{[]int{1, -2, 10, -5}}, []int{-2, 10}},
		{"tc5", args{[]int{-4, 6}}, []int{-4, 6}},
		{"tc6", args{[]int{4, -6}}, []int{-6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
