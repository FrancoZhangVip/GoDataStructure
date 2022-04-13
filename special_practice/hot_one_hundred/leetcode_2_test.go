package hot_one_hundred

import (
	"fmt"
	"testing"
)

func ShowListNodes(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, "\t")
	}
	fmt.Println()
}

func TestAddTwoNumbers(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)

	node1.Val = 2
	node2.Val = 4
	node3.Val = 3

	node1.Next = node2
	node2.Next = node3

	node4 := new(ListNode)
	node5 := new(ListNode)
	node6 := new(ListNode)

	node4.Val = 5
	node5.Val = 6
	node6.Val = 4

	node4.Next = node5
	node5.Next = node6

	res := AddTwoNumbers(node1, node4)
	ShowListNodes(res)
}

func TestAddTwoNumbers2(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 0

	node4 := new(ListNode)
	node4.Val = 0

	res := AddTwoNumbers(node1, node4)
	ShowListNodes(res)
}

func TestAddTwoNumbers3(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node8 := new(ListNode)
	node9 := new(ListNode)
	node10 := new(ListNode)
	node11 := new(ListNode)

	node1.Val = 9
	node2.Val = 9
	node3.Val = 9
	node8.Val = 9
	node9.Val = 9
	node10.Val = 9
	node11.Val = 9

	node1.Next = node2
	node2.Next = node3
	node3.Next = node8
	node8.Next = node9
	node9.Next = node10
	node10.Next = node11

	node4 := new(ListNode)
	node5 := new(ListNode)
	node6 := new(ListNode)
	node7 := new(ListNode)

	node4.Val = 9
	node5.Val = 9
	node6.Val = 9
	node7.Val = 9

	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	res := AddTwoNumbers(node1, node4)
	ShowListNodes(res)
}
