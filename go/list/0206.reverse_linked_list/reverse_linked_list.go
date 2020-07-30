package _206_reverse_linked_list

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

type ListNode struct {
	Val  int
	Next *ListNode
}

//         1 -> 2 -> 3 -> 4 -> 5 -> null
//null  <- 1 <- 2 <- 3 <- 4 <- 5
func reverseList(head *ListNode) *ListNode {
	var newHead, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}
