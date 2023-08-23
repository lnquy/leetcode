package main

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name      string
		args      args
		wantFlips int
	}{
		{"tc1", args{2, 6, 5}, 3},
		{"tc2", args{4, 2, 7}, 1},
		{"tc3", args{1, 2, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFlips := minFlips(tt.args.a, tt.args.b, tt.args.c); gotFlips != tt.wantFlips {
				t.Errorf("minFlips() = %v, want %v", gotFlips, tt.wantFlips)
			}
		})
	}
}
