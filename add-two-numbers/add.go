package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
	cry := 0
	head := (*ListNode)(nil)
	tail := (*ListNode)(nil)

	for list1 != nil || list2 != nil || cry > 0 {
		val := cry

		if list1 != nil {
			val += list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			val += list2.Val
			list2 = list2.Next
		}

		cry = val / 10
		val = val % 10

		node := &ListNode{Val: val, Next: nil}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return head
}
