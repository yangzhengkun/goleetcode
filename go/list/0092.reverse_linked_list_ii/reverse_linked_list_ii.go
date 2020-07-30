package _092_reverse_linked_list_ii

//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
//说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
//示例:
//
//输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := &ListNode{Next: head}
	prem := dummy
	for i := 0; i < m; i++ {
		prem = prem.Next
	}
	cur, next := prem.Next, prem.Next.Next
	var pre, nextnext *ListNode
	for i := m; i < n; i++ {
		nextnext = next.Next
		cur.Next, pre, next.Next = pre, cur, cur
		cur, next = next, nextnext
	}
	prem.Next.Next = next
	prem.Next = cur
	return dummy.Next
}
