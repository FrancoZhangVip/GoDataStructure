package hot_one_hundred

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestHasCycle(t *testing.T) {
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

	assert.Equal(t, true, HasCycle(node1))
}
