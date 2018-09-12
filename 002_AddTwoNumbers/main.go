package main

import "fmt"

func main() {
	var l1, l2 *ListNode
	l1 = insertTail(l1, 2)
	l1 = insertTail(l1, 4)
	l1 = insertTail(l1, 3)

	l2 = insertTail(l2, 5)
	l2 = insertTail(l2, 6)
	l2 = insertTail(l2, 4)

	fmt.Println(dumpList(addTwoNumbers(l1, l2)))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, n2 := l1, l2
	var out *ListNode

	sum, dozens := 0, 0
	for n1 != nil || n2 != nil {
		sum, dozens = dozens, 0
		if n1 != nil {
			sum += n1.Val
		}
		if n2 != nil {
			sum += n2.Val
		}
		if sum > 9 {
			dozens = sum / 10
			sum %= 10
		}
		out = insertTail(out, sum)

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}

	if dozens != 0 {
		out = insertTail(out, dozens)
	}
	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertTail(l *ListNode, v int) *ListNode {
	if l == nil {
		return &ListNode{Val: v}
	}
	n := l
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &ListNode{Val: v}
	return l
}

func dumpList(l *ListNode) []int {
	var out []int
	n := l
	for n != nil {
		out = append(out, n.Val)
		n = n.Next
	}
	return out
}
