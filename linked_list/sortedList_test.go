package linked_list

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMid(head)
	l1 := sortList(head)
	l2 := sortList(mid)

	return merge(l1, l2)
}

func findMid(head *ListNode) *ListNode {
	var slow, fast, tmp *ListNode
	slow = head
	fast = head

	for fast != nil && fast.Next != nil {
		tmp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp.Next = nil
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l2.Next, l1)
		return l2
	}

}

//快速排序
func sortList(head *ListNode) *ListNode {
	var start, end *ListNode
	start, end = head, nil
	partition(start, end)
	return start
}

func partition(start, end *ListNode) {
	if start == end || start.Next == end {
		return
	}
	base, target := start, start

	for i := start.Next; i != end; i = i.Next {
		if i.Val < base.Val {
			target = target.Next
			target.Val, i.Val = i.Val, target.Val
		}
	}

	base.Val, target.Val = target.Val, base.Val
	partition(start, target)
	partition(target.Next, end)
}

func TestSortedLinkedList(t *testing.T) {

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

	for {
		fmt.Print(p.Val, "->")
		p = p.Next
		if p.Next == nil {
			fmt.Print(p.Val)
			break
		}
	}
	fmt.Println()
	p = sortList(&l1)
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
