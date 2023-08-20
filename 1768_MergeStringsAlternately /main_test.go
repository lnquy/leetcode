package main

import "testing"

func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"tc1", args{"abc", "pqr"}, "apbqcr"},
		{"tc2", args{"ab", "pqrs"}, "apbqrs"},
		{"tc3", args{"abcd", "pq"}, "apbqcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
