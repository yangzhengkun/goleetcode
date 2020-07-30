package _206_reverse_linked_list

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	var head = &ListNode{
		Val:  1,
		Next: nil,
	}

	var node2 = &ListNode{Val: 2, Next: nil}
	var node3 = &ListNode{Val: 3, Next: nil}
	var node4 = &ListNode{Val: 4, Next: nil}
	var node5 = &ListNode{Val: 5, Next: nil}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	ret := reverseList(head)
	for ret != nil {
		fmt.Printf("%d->", ret.Val)
		ret = ret.Next
	}
}
