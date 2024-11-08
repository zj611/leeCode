package linked_list

import (
	"fmt"
	"testing"
)

//实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
//输入： 1->2->3->4->5 和 k = 2
//输出： 4
//Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 快慢指针方法

func kthToLast(head *ListNode, k int) int {

	if head == nil {
		return 0
	}

	if head.Next == nil {
		return head.Val
	}

	var fast, slow *ListNode
	fast = head
	slow = head
	var i int
	for i = 0; i < k; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return slow.Val
	}

	for {
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Next.Val
}

func TestKthToLast(t *testing.T) {

	l1 := ListNode{
		Val: 11,
	}
	l2 := ListNode{
		Val: 22,
	}
	l3 := ListNode{
		Val: 3,
	}
	l4 := ListNode{
		Val: 14,
	}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = nil

	p := &l1
	PrintLinkedList(p)

	res := kthToLast(&l1, 2)

	fmt.Println(res)

}
