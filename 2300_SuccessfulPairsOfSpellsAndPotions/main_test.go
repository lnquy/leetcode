package main

import (
	"reflect"
	"testing"
)

func Test_successfulPairs(t *testing.T) {
	type args struct {
		spells  []int
		potions []int
		success int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{"tc1", args{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7}, []int{4, 0, 3}},
		{"tc2", args{[]int{3, 1, 2}, []int{8, 5, 8}, 16}, []int{2, 0, 2}},
		{"tc3", args{[]int{15, 8, 19}, []int{38, 36, 23}, 328}, []int{3, 0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := successfulPairs(tt.args.spells, tt.args.potions, tt.args.success); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("successfulPairs() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
