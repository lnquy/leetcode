package main

import "fmt"

func main() {
	var l1, t1 *ListNode
	l1, t1 = insertTail(l1, t1, &ListNode{Val: 1})
	l1, t1 = insertTail(l1, t1, &ListNode{Val: 2})
	l1, t1 = insertTail(l1, t1, &ListNode{Val: 4})
	var l2, t2 *ListNode
	l2, t2 = insertTail(l2, t2, &ListNode{Val: 1})
	l2, t2 = insertTail(l2, t2, &ListNode{Val: 3})
	l2, t2 = insertTail(l2, t2, &ListNode{Val: 4})

	out := mergeTwoLists(l1, l2)
	fmt.Println(dumpList(out))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(m+n)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	n1, n2 := l1, l2
	var head, tail *ListNode

	for n1 != nil || n2 != nil {
		var tmp *ListNode

		if n1 != nil && n2 != nil {
			if n1.Val < n2.Val {
				tmp, n1 = pop(n1)
				head, tail = insertTail(head, tail, tmp)
				continue
			}
			tmp, n2 = pop(n2)
			head, tail = insertTail(head, tail, tmp)
			continue
		}

		if n1 == nil {
			head, tail = insertTail(head, tail, n2)
			break
		}
		head, tail = insertTail(head, tail, n1)
		break
	}

	return head
}

func pop(l *ListNode) (*ListNode, *ListNode) {
	if l == nil {
		return nil, nil
	}
	carry := l.Next
	l.Next = nil
	return l, carry
}

func insertTail(head, tail, n *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return n, n
	}
	tail.Next = n
	return head, n
}

// Helpers
func dumpList(l *ListNode) []int {
	var out []int
	n := l
	for n != nil {
		out = append(out, n.Val)
		n = n.Next
	}
	return out
}
