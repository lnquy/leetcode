package main

import "fmt"

func main() {
	in := arrayToLinkedList([]int{0, 1, 2})
	fmt.Println("IN", linkedListToArray(in))
	g := deleteMiddle(in)
	fmt.Println("OUT", linkedListToArray(g))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	next := head.Next
	if next == nil { // len = 1
		return nil
	}
	next = next.Next
	if next == nil { // len = 2
		head.Next = nil
		return head
	}
	next = next.Next
	if next == nil { // len = 3
		head.Next = head.Next.Next
		return head
	}

	slowPtr := head.Next
	counter := 0
	for next.Next != nil { // len >= 3
		next = next.Next
		counter++
		// Only move slow ptr once every 2 node traversed
		// So when we reached the end of the linked list, the slowPtr is right before the
		// middle node that need to be removed
		if counter != 0 && counter%2 == 0 {
			slowPtr = slowPtr.Next
		}
	}

	slowPtr.Next = slowPtr.Next.Next // Remove middle
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
