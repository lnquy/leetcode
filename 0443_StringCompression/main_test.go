package main

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"tc1", args{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}}, 6},
		{"tc2", args{[]byte{'a'}}, 1},
		{"tc3", args{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}}, 4},
		{"tc4", args{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
