package list

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)

	node1.Val = 4
	node2.Val = 1
	node3.Val = 8
	node4.Val = 4
	node5.Val = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	nodea := new(ListNode)
	nodeb := new(ListNode)
	nodec := new(ListNode)

	nodea.Val = 5
	nodeb.Val = 0
	nodec.Val = 1

	nodea.Next = nodeb
	nodeb.Next = nodec
	nodec.Next = node3

	ShowList(node1)
	ShowList(nodea)

	println(GetIntersectionNode(node1, nodea).Val)
}

func TestGetIntersectionNode2(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)

	node1.Val = 0
	node2.Val = 9
	node3.Val = 1
	node4.Val = 2
	node5.Val = 4

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	nodea := new(ListNode)
	nodea.Val = 3

	nodea.Next = node4

	ShowList(node1)
	ShowList(nodea)
	println(GetIntersectionNode(node1, nodea).Val)
}

func TestGetIntersectionNode3(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)

	node1.Val = 2
	node2.Val = 6
	node3.Val = 4

	node1.Next = node2
	node2.Next = node3

	nodea := new(ListNode)
	nodeb := new(ListNode)
	nodea.Val = 1
	nodeb.Val = 5

	ShowList(node1)
	ShowList(nodea)
	println(GetIntersectionNode(node1, nodea) == nil)
}
