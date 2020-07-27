package _141_linked_list_cycle

//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//示例 1：
//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//示例 2：
//
//输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。

//示例 3：
//
//输入：head = [1], pos = -1
//输出：false
//解释：链表中没有环。

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		// 头结点为空或只有一个节点
		return false
	}
	var fast = head.Next
	var slow = head

	for fast != slow {
		if fast == nil || fast.Next == nil {
			// fast指针走到头
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
