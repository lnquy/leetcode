package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	headEvenNode := &ListNode{Next: head.Next}
	oddNode, evenNode := head, head.Next

	for evenNode != nil && evenNode.Next != nil {
		oddNode.Next = evenNode.Next
		oddNode = oddNode.Next
		evenNode.Next = oddNode.Next
		evenNode = evenNode.Next
	}

	oddNode.Next = headEvenNode.Next

	return head
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
