//
// Created by yangzhengkun on 2020/7/26.
//

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */


struct ListNode {
  int val;
  ListNode *next;
  explicit ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
 public:
  ListNode *removeNthFromEnd(ListNode *head, int n) {
	if (head == nullptr || head->next == nullptr) {
	  return nullptr;
	}
	ListNode dummyHead = ListNode(0);
	dummyHead.next = head;
	ListNode *fast = &dummyHead;
	ListNode *slow = &dummyHead;

	fast->next = head, slow->next = head;
	int cnt = 0;
	while (fast) {
	  fast = fast->next;
	  if (cnt > n) {
		slow = slow->next;
	  }
	  cnt++;
	}
	ListNode *rm = slow->next;
	slow->next = rm->next;
	delete rm;
	return dummyHead.next;
  }
};
