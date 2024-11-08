package linked_list

import "testing"

// 旋转链表k次
// Input: 0->1->2->NULL, k = 4
// Output: 2->0->1->NULL
// Explanation:
// rotate 1 steps to the right: 2->0->1->NULL
// rotate 2 steps to the right: 1->2->0->NULL
// rotate 3 steps to the right: 0->1->2->NULL
// rotate 4 steps to the right: 2->0->1->NULL
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	length := 0
	cur := newHead
	// 获取链表总长度
	for cur.Next != nil {
		length++
		cur = cur.Next
	}
	// 构成环
	cur.Next = head

	// 取余数，然后再移动指针
	if (k % length) == 0 {
		return head
	}

	// cur再跑到初始位置
	cur = newHead
	for i := length - k%length; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}

func rotateRight1(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil || n == 0 {
		return head
	}

	// 获取链表长度
	var length int
	newHead := &ListNode{
		Val: 0, Next: head,
	}
	cur := newHead

	for cur.Next != nil {
		length++
		cur = cur.Next
	}

	// 构成环
	cur.Next = head

	moveTimes := n % length
	if moveTimes == 0 {
		return head
	}

	cur = newHead

	for i := length - moveTimes; i > 0; i-- {
		cur = cur.Next
	}

	res := &ListNode{
		Val:  0,
		Next: cur.Next,
	}
	cur.Next = nil
	return res.Next

}

func TestRotateRight(t *testing.T) {

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
	p = rotateRight1(&l1, 3)
	PrintLinkedList(p)
}
