package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func arrayToList(nums ...int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for i := range nums {
		node := &ListNode{Val: nums[i], Next: nil}
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

func dumpList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d", list.Val)
		if list.Next != nil {
			fmt.Printf(" -> ")
		}
		list = list.Next
	}
	fmt.Printf("\n")
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode
	var tail *ListNode

	for list1 != nil && list2 != nil {
		var node *ListNode

		if list1.Val < list2.Val {
			node = list1
			list1 = list1.Next
		} else {
			node = list2
			list2 = list2.Next
		}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
		tail.Next = nil
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return mergeTwoList(lists[0], lists[1])
	}

	mid := len(lists) / 2
	l1 := mergeKLists(lists[:mid])
	l2 := mergeKLists(lists[mid:])
	return mergeTwoList(l1, l2)
}

func mergeKLists1(lists []*ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode

	for {
		sel := int(-1)
		min := int(math.MaxInt32)
		for i := range lists {
			if lists[i] != nil {
				if lists[i].Val < min {
					sel = i
					min = lists[i].Val
				}
			}
		}

		if sel < 0 {
			break
		} else {
			lists[sel] = lists[sel].Next
		}

		node := &ListNode{Val: min, Next: nil}
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

func main() {
	list1 := arrayToList(1, 4, 5)
	list2 := arrayToList(1, 3, 4)
	list3 := mergeTwoList(list1, list2)
	dumpList(list3)

	var lists []*ListNode
	lists = append(lists, arrayToList(1, 4, 5))
	lists = append(lists, arrayToList(1, 3, 4))
	lists = append(lists, arrayToList(2, 6))
	merged := mergeKLists(lists)
	dumpList(merged)
}
