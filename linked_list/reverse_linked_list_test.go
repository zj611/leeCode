package linked_list

import (
	"fmt"
	"testing"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

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

	for {
		fmt.Print(p.Val, "->")
		p = p.Next
		if p.Next == nil {
			fmt.Print(p.Val)
			break
		}
	}
	fmt.Println()
	p = ReverseLinkedList(&l1)
	for {
		fmt.Print(p.Next, "->")
		p = p.Next
		if p.Next == nil {
			fmt.Print(p.Val)
			break
		}
	}
	fmt.Println()

}
