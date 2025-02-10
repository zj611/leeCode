package linked_list

import (
	"testing"
)

// 删除链表中重复的结点，只要是有重复过的结点，只保留一个
// 注意：存在问题，当重复结点不在前后位置时，容易漏掉，只能是有序链表
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func TestDeleteDuplicates(t *testing.T) {

	l1 := ListNode{
		Val: 1,
	}
	l2 := ListNode{
		Val: 1,
	}
	l3 := ListNode{
		Val: 3,
	}
	l4 := ListNode{
		Val: 1,
	}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = nil

	p := &l1
	PrintLinkedList(p)
	p = deleteDuplicates(&l1)
	PrintLinkedList(p)
}
