package main

import (
	"reflect"
	"testing"
)

var tcs = []struct {
	inNum int
	out   [][]int
}{
	{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},

	{0, [][]int{}},
	{1, [][]int{{1}}},
	{2, [][]int{{1}, {1, 1}}},
	{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
	{4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
	{6, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}}},
	{7, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}}},
	{8, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}, {1, 7, 21, 35, 35, 21, 7, 1}}},
	{9, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}, {1, 7, 21, 35, 35, 21, 7, 1}, {1, 8, 28, 56, 70, 56, 28, 8, 1}}},
	{10, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}, {1, 7, 21, 35, 35, 21, 7, 1}, {1, 8, 28, 56, 70, 56, 28, 8, 1}, {1, 9, 36, 84, 126, 126, 84, 36, 9, 1}}},
}

func TestGenerate(t *testing.T) {
	for idx, tc := range tcs {
		out := generate(tc.inNum)
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("%d. generate(%d) = %v. Expected: %v", idx, tc.inNum, out, tc.out)
		}
	}
}

func TestGenerate_big(t *testing.T) {
	for idx, tc := range tcs {
		out := generate_big(tc.inNum)
		if !reflect.DeepEqual(out, tc.out) {
			t.Errorf("%d. generate_big(%d) = %v. Expected: %v", idx, tc.inNum, out, tc.out)
		}
	}
}

var _dump [][]int

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_dump = generate(1000)
	}
}

func BenchmarkGenerate_big(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_dump = generate_big(1000)
	}
}
