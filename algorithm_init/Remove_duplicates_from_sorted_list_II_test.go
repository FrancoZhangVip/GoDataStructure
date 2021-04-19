package algorithm_init

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicatesII(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 3
	node5 := new(ListNode)
	node5.Val = 4
	node6 := new(ListNode)
	node6.Val = 4
	node7 := new(ListNode)
	node7.Val = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	res := DeleteDuplicatesII(node1)

	for res != nil {
		fmt.Print(res.Val, "\t")
		res = res.Next
	}
}
