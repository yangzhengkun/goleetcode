package _019_remove_nth_node_from_end_of_list

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var dummyHead = ListNode{Val: 0}
	dummyHead.Next = head
	var fast = &dummyHead
	var slow = &dummyHead

	var cnt = 0
	for fast != nil {
		fast = fast.Next
		if cnt > n {
			slow = slow.Next
		}
		cnt += 1
	}
	//slow最终指向的是 待删除节点的前一个节点
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
