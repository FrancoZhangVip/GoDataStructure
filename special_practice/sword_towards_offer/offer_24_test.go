package sword_towards_offer

import (
	"fmt"
	"testing"
)

func ShowList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, "\t")
	}
	fmt.Println()
}

func TestReverseList(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1

	node2 := new(ListNode)
	node2.Val = 2

	node3 := new(ListNode)
	node3.Val = 3

	node4 := new(ListNode)
	node4.Val = 4

	node5 := new(ListNode)
	node5.Val = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	ShowList(ReverseList(node1))
}
