package linked_list

import (
	"fmt"
	"testing"
)

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

// 按奇偶数分割链表
func splite(head *ListNode2) *ListNode2 {

	if head == nil || head.Next == nil {
		return nil
	}

	p := head
	q := head.Next

	M := &ListNode2{Val: 0}
	M.Next = q

	for q != nil {
		p.Next = q.Next
		p = q
		q = q.Next
	}
	p.Next = nil
	return M.Next
}

func TestSplite(t *testing.T) {

	n1 := &ListNode2{Val: 1}
	n2 := &ListNode2{Val: 2}
	n3 := &ListNode2{Val: 3}
	n4 := &ListNode2{Val: 4}
	n5 := &ListNode2{Val: 5}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	head := n1

	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println()

	m := splite(n1)

	head = n1
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println()

	head = m
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println()
}
