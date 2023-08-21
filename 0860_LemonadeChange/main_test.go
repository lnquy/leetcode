package main

import "testing"

func TestLemonadeChange(t *testing.T) {
	tcs := []struct {
		inBills []int
		out     bool
	}{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
		{[]int{5, 5, 10, 10, 20}, false},

		{[]int{}, true},
		{[]int{5}, true},
		{[]int{5, 5, 5, 5, 5, 5}, true},
		{[]int{10}, false},
		{[]int{5, 10}, true},
		{[]int{5, 10, 10}, false},
		{[]int{5, 10, 5, 20}, true},
		{[]int{5, 10, 5, 20}, true},
		{[]int{5, 5, 5, 5, 20, 10}, true},
		{[]int{5, 5, 5, 5, 20, 10, 10}, false},
		{[]int{5, 5, 5, 5, 20, 10, 5, 10}, true},
		{[]int{5, 5, 5, 5, 20, 10, 20}, false},
	}

	for idx, tc := range tcs {
		out := lemonadeChange(tc.inBills)
		if out != tc.out {
			t.Errorf("%d. lemonadeChange(%v) = %t. Expected: %t", idx, tc.inBills, out, tc.out)
		}
	}

}
