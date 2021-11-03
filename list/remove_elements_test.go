package list

import "testing"

func TestRemoveElements(t *testing.T) {
	//head = [1,2,6,3,4,5,6], val = 6
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)
	node6 := new(ListNode)
	node7 := new(ListNode)

	node1.Val = 1
	node2.Val = 2
	node3.Val = 6
	node4.Val = 3
	node5.Val = 4
	node6.Val = 5
	node7.Val = 6

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	ShowList(node1)
	head := RemoveElements(node1, 6)
	ShowList(head)
}

func TestRemoveElements2(t *testing.T) {

	head := RemoveElements(nil, 1)
	ShowList(head)
}

func TestRemoveElements3(t *testing.T) {

	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)

	node1.Val = 7
	node2.Val = 7
	node3.Val = 7
	node4.Val = 7

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	ShowList(node1)
	head := RemoveElements(node1, 7)
	ShowList(head)
}

func ShowList(head *ListNode) {
	for head != nil {
		print(head.Val, "->")
		head = head.Next
	}
	println()
}
