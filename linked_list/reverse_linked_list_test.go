package linked_list

import (
	"fmt"
	"testing"
)

func PrintLinkedList(p *ListNode) {
	for {
		fmt.Print(p.Val, "->")
		p = p.Next
		if p.Next == nil {
			fmt.Print(p.Val)
			break
		}
	}
	fmt.Println()
}

type LinkedNode1 struct {
	val  int
	next *LinkedNode1
}

func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, p, next *ListNode
	pre = nil
	p = head
	next = head.Next
	for {
		if p == nil {
			break
		}
		p.Next = pre
		pre = p
		p = next
		if next != nil {
			next = next.Next
		}

	}
	return pre
}

func ReverseLinkedList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, p, next *ListNode
	pre = nil
	p = head
	next = head.Next
	for p != nil {
		p.Next = pre
		pre = p
		p = next
		if next != nil {
			next = next.Next
		}
	}
	return pre
}

func TestReverseLinkedList(t *testing.T) {

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

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = nil

	p := &l1
	PrintLinkedList(p)
	p = ReverseLinkedList1(&l1)
	PrintLinkedList(p)
}
