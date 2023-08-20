package main

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"tc1", args{[]int{1, 0, 0, 0, 1}, 1}, true},
		{"tc2", args{[]int{1, 0, 0, 0, 1}, 2}, false},
		{"tc3", args{[]int{1, 0, 1, 0, 1}, 1}, false},
		{"tc4", args{[]int{0, 0, 0}, 2}, true},
		{"tc5", args{[]int{0, 1, 0, 0, 0}, 1}, true},
		{"tc6", args{[]int{0, 1, 0, 0, 1}, 1}, false},
		{"tc7", args{[]int{0, 1, 0, 1, 0}, 1}, false},
		{"tc8", args{[]int{0, 1, 0, 1, 1}, 1}, false},
		{"tc9", args{[]int{0, 1, 1, 0, 0}, 1}, true},
		{"tc10", args{[]int{0, 1, 1, 0, 1}, 1}, false},
		{"tc11", args{[]int{0, 1, 1, 1, 0}, 1}, false},
		{"tc12", args{[]int{0, 1, 1, 1, 1}, 1}, false},
		{"tc13", args{[]int{1, 0, 0, 0, 0, 1}, 2}, false},
		// 1, 0, 0, 0, 0, 0, 1 => 1, 0, [1], 0, [1], 0, 1
		// 0, 0, 0, 0, 0, 0 => [1], 0, [1], 0, [1], 0
		// 0, 0, 0, 0, 0, 0, 0 => [1], 0, [1], 0, [1], 0, [1]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
