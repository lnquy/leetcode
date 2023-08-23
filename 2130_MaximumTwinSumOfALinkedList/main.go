package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) (max int) {
	arr := make([]int, 0, 8)
	n := head
	for n != nil {
		arr = append(arr, n.Val)
		n = n.Next
	}

	var sum int
	last := len(arr) - 1
	for i := 0; i < len(arr)/2; i++ {
		sum = arr[last-i] + arr[i]
		if max < sum {
			max = sum
		}
	}
	return max
}

// ---------
// Helper functions, has no affect to the correctness of the solution

func (l *ListNode) Push(v int) {
	next := l
	for next.Next != nil {
		next = next.Next
	}
	next.Next = &ListNode{
		Val:  v,
		Next: nil,
	}
}

func arrayToLinkedList(arr []int) *ListNode {
	if len(arr) <= 0 {
		return nil
	}

	head := ListNode{
		Val:  arr[0],
		Next: nil,
	}

	for i := 1; i < len(arr); i++ {
		head.Push(arr[i])
	}
	return &head
}

func linkedListToArray(head *ListNode) []int {
	arr := make([]int, 0)
	n := head
	for n != nil {
		arr = append(arr, n.Val)
		n = n.Next
	}
	return arr
}
