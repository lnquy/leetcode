package main

import "testing"

func Test_equalPairs(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"tc1", args{[][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}}, 1},
		{"tc2", args{[][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}}, 3},
		{"tc3", args{[][]int{{13, 13}, {13, 13}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalPairs(tt.args.grid); got != tt.want {
				t.Errorf("equalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

var _res int

func BenchmarkEqualPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_res = equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}})
	}
}
