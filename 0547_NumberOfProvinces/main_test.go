package main

import "testing"

func Test_findCircleNum(t *testing.T) {
	type args struct {
		isConnected [][]int
	}
	tests := []struct {
		name          string
		args          args
		wantProvinces int
	}{
		{"tc1", args{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}}, 2},
		{"tc2", args{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProvinces := findCircleNum(tt.args.isConnected); gotProvinces != tt.wantProvinces {
				t.Errorf("findCircleNum() = %v, want %v", gotProvinces, tt.wantProvinces)
			}
		})
	}
}
