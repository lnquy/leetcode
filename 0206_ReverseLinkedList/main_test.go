package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"tc1", args{[]int{1, 2, 3, 4, 5}}, []int{5, 4, 3, 2, 1}},
		{"tc2", args{[]int{1, 2}}, []int{2, 1}},
		{"tc3", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(arrayToLinkedList(tt.args.head)); !reflect.DeepEqual(linkedListToArray(got), tt.want) {
				t.Errorf("reverseList() = %v, want %v", linkedListToArray(got), tt.want)
			}
		})
	}
}
