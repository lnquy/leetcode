package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head.Next
	slow := head
	var tmp *ListNode
	slow.Next = nil
	for fast != nil {
		tmp = fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}
	return slow
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
