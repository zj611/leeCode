package linked_list

import "testing"

// 给定一个数 x，比 x 大或等于的数字都要排列在比 x 小的数字后面，并且相对位置不能发生变化。
// 由于相对位置不 能发生变化，所以不能用类似冒泡排序的思想。
// 解法一 单链表
func partitionLinkedList(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{Val: 0, Next: nil}
	before := beforeHead
	afterHead := &ListNode{Val: 0, Next: nil}
	after := afterHead
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}

func partitionLinkedList1(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{
		Val: 0,
	}

	afterHead := &ListNode{
		Val: 0,
	}

	before := beforeHead
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}

	after.Next = nil
	before.Next = afterHead.Next

	return beforeHead.Next

}

func TestPartitionLinkedList(t *testing.T) {

	l1 := ListNode{
		Val: 1,
	}
	l2 := ListNode{
		Val: 2,
	}
	l3 := ListNode{
		Val: 4,
	}
	l4 := ListNode{
		Val: 3,
	}
	l5 := ListNode{
		Val: 7,
	}
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = nil

	p := &l1
	PrintLinkedList(p)

	p = partitionLinkedList1(&l1, 4)
	PrintLinkedList(p)
}
