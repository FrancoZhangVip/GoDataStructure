package sword_towards_offer

import (
	"fmt"
	"testing"
)

func ShowNums(nums []int) {
	for _, data := range nums {
		fmt.Print(data, "\t")
	}
	fmt.Println()
}

func TestReversePrint(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1

	node2 := new(ListNode)
	node2.Val = 3

	node3 := new(ListNode)
	node3.Val = 2

	node1.Next = node2
	node2.Next = node3

	ShowNums(ReversePrint(node1))
}
