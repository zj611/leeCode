package linked_list

import (
	"fmt"
	"testing"
)

type LinkedList struct {
	 value  int
	 nextPtr   *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList  {
	if head == nil || head.nextPtr == nil{
		return head
	}

	var pre,p,next *LinkedList
	pre = nil
	p = head
	next = head.nextPtr
	for{
		if p == nil{
			break
		}
		p.nextPtr = pre
		pre = p
		p = next
		if next != nil{
			next = next.nextPtr
		}
	}

	return pre

}

func TestReverseLinkedList(t *testing.T)  {

	l1 := LinkedList{
		value: 1,
	}
	l2 := LinkedList{
		value: 2,
	}
	l3 := LinkedList{
		value: 3,
	}
	l4 := LinkedList{
		value: 4,
	}

	l1.nextPtr = &l2
	l2.nextPtr = &l3
	l3.nextPtr = &l4
	l4.nextPtr = nil

	p := &l1

	for{
		fmt.Print(p.value, "->")
		p = p.nextPtr
		if p.nextPtr == nil{
			fmt.Print(p.value)
			break
		}
	}
	fmt.Println()
	p = ReverseLinkedList(&l1)
	for{
		fmt.Print(p.value, "->")
		p = p.nextPtr
		if p.nextPtr == nil{
			fmt.Print(p.value)
			break
		}
	}
	fmt.Println()

}

