package main

import (
	"container/heap"
)

type SmallestInfiniteSet struct {
	minHeap  *IntHeap
	itemSet  map[int]struct{}
	smallest int
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{
		minHeap:  h,
		itemSet:  make(map[int]struct{}),
		smallest: 0,
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.minHeap.Len() <= 0 {
		s.smallest++
		return s.smallest
	}
	m := heap.Pop(s.minHeap).(int)
	delete(s.itemSet, m)
	return m
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num <= s.smallest {
		if _, ok := s.itemSet[num]; !ok {
			heap.Push(s.minHeap, num)
			s.itemSet[num] = struct{}{}
		}
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
