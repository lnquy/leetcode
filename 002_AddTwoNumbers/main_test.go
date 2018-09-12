package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tcs := []struct {
		inNum1 []int
		inNum2 []int
		outSum []int
	}{
		{inNum1: []int{0}, inNum2: []int{0}, outSum: []int{0}},
		{inNum1: []int{1}, inNum2: []int{1}, outSum: []int{2}},
		{inNum1: []int{5}, inNum2: []int{5}, outSum: []int{0, 1}},
		{inNum1: []int{0, 1}, inNum2: []int{0, 2}, outSum: []int{0, 3}},
		{inNum1: []int{2, 4, 3}, inNum2: []int{5, 6, 4}, outSum: []int{7, 0, 8}},
		{inNum1: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, inNum2: []int{5, 6, 4}, outSum: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for idx, tc := range tcs {
		l1, l2 := insertBatch(tc.inNum1), insertBatch(tc.inNum2)
		out := dumpList(addTwoNumbers(l1, l2))
		if !reflect.DeepEqual(out, tc.outSum) {
			t.Errorf("%d. addTwoNumbers(%v, %v) = %v. Expected: %v", idx, tc.inNum1, tc.inNum2, out, tc.outSum)
		}
	}
}

func insertBatch(nums []int) *ListNode {
	var l *ListNode

	for _, v := range nums {
		l = insertTail(l, v)
	}
	return l
}
