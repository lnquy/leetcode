package main

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"tc1", args{[]int{1, 2, 3, 4, 5}}, []int{1, 3, 5, 2, 4}},
		{"tc1", args{[]int{2, 1, 3, 5, 6, 4, 7}}, []int{2, 3, 6, 7, 1, 5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(arrayToLinkedList(tt.args.head)); !reflect.DeepEqual(linkedListToArray(got), tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", linkedListToArray(got), tt.want)
			}
		})
	}
}
