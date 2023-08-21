package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tcs := []struct {
		inList1 []int
		inList2 []int
		outList []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},

		{[]int{}, []int{}, nil},
		{[]int{1}, []int{}, []int{1}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{2}, []int{2}},
		{[]int{}, []int{2, 4, 5, 6}, []int{2, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{1, 1, 1, 2, 2, 3, 5}, []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 5}},
		{[]int{1, 3, 5, 7}, []int{2, 4, 6, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for idx, tc := range tcs {
		l1, l2 := insertBatch(tc.inList1), insertBatch(tc.inList2)
		out := dumpList(mergeTwoLists(l1, l2))

		if !reflect.DeepEqual(out, tc.outList) {
			t.Errorf("%d. mergeTwoLists(%v, %v) = %v. Expected: %v", idx, tc.inList1, tc.inList2, out, tc.outList)
		}
	}
}

func insertBatch(nums []int) *ListNode {
	var h, t *ListNode

	for _, v := range nums {
		h, t = insertTail(h, t, &ListNode{Val: v})
	}
	return h
}
