package list

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)

	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	ShowList(node1)

	head := RemoveNthFromEnd(node1, 2)

	ShowList(head)
}

func TestRemoveNthFromEnd2(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1

	ShowList(node1)

	head := RemoveNthFromEnd(node1, 1)

	ShowList(head)
}

func TestRemoveNthFromEnd3(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)

	node1.Val = 1
	node2.Val = 2

	node1.Next = node2

	ShowList(node1)

	head := RemoveNthFromEnd(node1, 1)

	ShowList(head)
}

func TestRemoveNthFromEnd4(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)

	node1.Val = 1
	node2.Val = 2

	node1.Next = node2

	ShowList(node1)

	head := RemoveNthFromEnd(node1, 2)

	ShowList(head)
}
