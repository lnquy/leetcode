package main

import (
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"tc1", args{nums1: []int{1, 2, 3}, nums2: []int{2, 4, 6}}, [][]int{{1, 3}, {4, 6}}},
		{"tc2", args{nums1: []int{1, 2, 3, 3}, nums2: []int{1, 1, 2, 2}}, [][]int{{3}, {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.args.nums1, tt.args.nums2)
			if len(got) != len(tt.want) && len(got) != 2 {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
				return
			}
			if !equalSlice(got[0], tt.want[0]) || !equalSlice(got[1], tt.want[1]) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]struct{}, len(a))
	for _, x := range a {
		m[x] = struct{}{}
	}

	for _, x := range b {
		if _, ok := m[x]; !ok {
			return false
		}
	}
	return true
}
