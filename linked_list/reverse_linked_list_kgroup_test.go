package linked_list

import (
	"testing"
)

// 按照每 K 个元素翻转的方式翻转链表。如果不满足 K 个元素的就不翻转

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)

	// 注意画图，reverse翻转前，head和node相差k个节点，reverse翻转后，head->next=node，注意画图
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}
func TestReverseLinkedListKGroup(t *testing.T) {

	l1 := ListNode{
		Val: 1,
	}
	l2 := ListNode{
		Val: 2,
	}
	l3 := ListNode{
		Val: 3,
	}
	l4 := ListNode{
		Val: 4,
	}
	l5 := ListNode{
		Val: 5,
	}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = nil

	p := &l1
	PrintLinkedList(p)
	p = reverseKGroup(&l1, 3)
	PrintLinkedList(p)
}
