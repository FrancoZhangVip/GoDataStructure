package algorithm_init

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 1
	node3 := new(ListNode)
	node3.Val = 2

	node1.Next = node2
	node2.Next = node3

	res := DeleteDuplicates(node1)

	for res != nil {
		fmt.Print(res.Val, "\t")
		res = res.Next
	}
}
