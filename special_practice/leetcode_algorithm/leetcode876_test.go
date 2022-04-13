package leetcode_algorithm

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestMiddleNode(t *testing.T) {
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

	assert.Equal(t, 3, MiddleNode(node1).Val)
}

func TestMiddleNode2(t *testing.T) {
	node1 := new(ListNode)
	node2 := new(ListNode)
	node3 := new(ListNode)
	node4 := new(ListNode)
	node5 := new(ListNode)
	node6 := new(ListNode)

	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5
	node6.Val = 6

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	assert.Equal(t, 4, MiddleNode(node1).Val)
}
