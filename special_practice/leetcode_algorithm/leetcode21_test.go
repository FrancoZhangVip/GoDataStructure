package leetcode_algorithm

import "testing"

func TestMergeTwoLists(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)
	node6 := new(ListNode)

	node1.Val = 1
	node2.Val = 2
	node3.Val = 4

	node1.Next = node2
	node2.Next = node3

	node4.Val = 1
	node5.Val = 3
	node6.Val = 4

	node4.Next = node5
	node5.Next = node6

	ShowList(MergeTwoLists(node1, node4))
}

func TestMergeTwoLists2(t *testing.T) {
	ShowList(MergeTwoLists(nil, nil))
}

func TestMergeTwoLists3(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 0
	ShowList(MergeTwoLists(nil, node1))
}
