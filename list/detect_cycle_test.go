package list

import "testing"

func TestDetectCycle(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)

	node1.Val = 3
	node2.Val = 2
	node3.Val = 0
	node4.Val = 4

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	node := DetectCycle(node1)
	if node != nil {
		println(node.Val)
	} else {
		println("nil")
	}

}

func TestDetectCycle2(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)

	node1.Val = 1
	node2.Val = 2

	node1.Next = node2
	node2.Next = node1

	node := DetectCycle(node1)
	if node != nil {
		println(node.Val)
	} else {
		println("nil")
	}

}

func TestDetectCycle3(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1

	node := DetectCycle(node1)
	if node != nil {
		println(node.Val)
	} else {
		println("nil")
	}

}
