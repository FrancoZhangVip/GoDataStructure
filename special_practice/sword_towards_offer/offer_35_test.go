package sword_towards_offer

import (
	"fmt"
	"testing"
)

func ShowRandomList(head *Node) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, "\t")
	}
	fmt.Println()
}

func ShowRandomVal(head *Node) {
	for p := head; p != nil; p = p.Next {
		if p.Random == nil {
			fmt.Print("nil", "\t")
		} else {
			fmt.Print(p.Val, "\t")
		}
	}
	fmt.Println()
}

func TestCopyRandomList(t *testing.T) {
	node1 := new(Node)
	node1.Val = 7

	node2 := new(Node)
	node2.Val = 13

	node3 := new(Node)
	node3.Val = 11

	node4 := new(Node)
	node4.Val = 10

	node5 := new(Node)
	node5.Val = 1

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node1.Random = nil
	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1

	res := CopyRandomList(node1)
	ShowRandomList(res)

	ShowRandomVal(node1)
	ShowRandomVal(res)
}
