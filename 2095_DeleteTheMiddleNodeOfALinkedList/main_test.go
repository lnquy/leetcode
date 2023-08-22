package main

import (
	"reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"tc1", args{[]int{1, 3, 4, 7, 1, 2, 6}}, []int{1, 3, 4, 1, 2, 6}},
		{"tc2", args{[]int{1, 2, 3, 4}}, []int{1, 2, 4}},
		{"tc3", args{[]int{2, 1}}, []int{2}},
		{"tc4", args{[]int{0}}, []int{}},
		{"tc5", args{[]int{0, 1}}, []int{0}},
		{"tc6", args{[]int{0, 1, 2}}, []int{0, 2}},
		{"tc7", args{[]int{0, 1, 2, 3}}, []int{0, 1, 3}},
		{"tc8", args{[]int{0, 1, 2, 3, 4}}, []int{0, 1, 3, 4}},
		{"tc9", args{[]int{0, 1, 2, 3, 4, 5}}, []int{0, 1, 2, 4, 5}},
		{"tc10", args{[]int{0, 1, 2, 3, 4, 5, 6}}, []int{0, 1, 2, 4, 5, 6}},
		{"tc11", args{[]int{0, 1, 2, 3, 4, 5, 6, 7}}, []int{0, 1, 2, 3, 5, 6, 7}},
		{"tc12", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}, []int{0, 1, 2, 3, 5, 6, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(arrayToLinkedList(tt.args.head)); !reflect.DeepEqual(linkedListToArray(got), tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", linkedListToArray(got), tt.want)
			}
		})
	}
}
