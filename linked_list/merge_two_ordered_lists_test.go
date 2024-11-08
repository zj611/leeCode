package linked_list

import (
	"testing"
)

/*
	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
	输入：l1 = [1,2,4], l2 = [1,3,4]
	输出：[1,1,2,3,4,4]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归的方法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func TestMergeTwoList(t *testing.T) {

	l1 := ListNode{
		Val: 1,
	}
	l2 := ListNode{
		Val: 2,
	}
	l3 := ListNode{
		Val: 4,
	}
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = nil

	m1 := ListNode{
		Val: 1,
	}
	m2 := ListNode{
		Val: 3,
	}
	m3 := ListNode{
		Val: 6,
	}
	m1.Next = &m2
	m2.Next = &m3
	m3.Next = nil

	p := &l1
	PrintLinkedList(p)
	p = &m1
	PrintLinkedList(p)

	p = mergeTwoLists(&l1, &m1)
	PrintLinkedList(p)
}
