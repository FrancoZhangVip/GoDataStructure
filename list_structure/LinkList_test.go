package list_structure

import (
	"fmt"
	"testing"
)

func TestLinkList_AddNode(t *testing.T) {
	node := NewListNode(5)
	node2 := NewListNode(25)
	node3 := NewListNode(35)
	list := NewLinkList()
	list.AddNode(node)
	fmt.Println("Size: ", list.Size())
	list.Show()
	list.AddNode(node2)
	fmt.Println("Size: ", list.Size())
	list.Show()
	list.AddNode(node3)
	fmt.Println("Size: ", list.Size())
	list.Show()
}

func TestLinkList_DelNode(t *testing.T) {
	node := NewListNode(5)
	node2 := NewListNode(25)
	node3 := NewListNode(35)
	list := NewLinkList()
	list.AddNode(node)
	list.AddNode(node2)
	list.AddNode(node3)

	list.DelNode(node2)
	fmt.Println("Size: ", list.Size())
	list.Show()
	fmt.Println("==================================")

	list = NewLinkList()
	list.AddNode(node)
	list.AddNode(node2)
	list.AddNode(node3)

	list.DelNode(node3)
	fmt.Println("Size: ", list.Size())
	list.Show()
	fmt.Println("==================================")

	list = NewLinkList()
	list.AddNode(node)
	list.AddNode(node2)
	list.AddNode(node3)

	list.DelNode(node)
	fmt.Println("Size: ", list.Size())
	list.Show()
	fmt.Println("==================================")

	list = NewLinkList()
	list.AddNode(node)
	list.AddNode(node2)
	list.AddNode(node3)

	list.DelNode(node)
	list.DelNode(node2)
	list.DelNode(node3)
	fmt.Println("Size: ", list.Size())
	list.Show()
	list.DelNode(node)
	fmt.Println("Size: ", list.Size())
	list.Show()
	fmt.Println("==================================")

	list = NewLinkList()
	list.AddNode(node)

	list.DelNode(node)
	fmt.Println("Size: ", list.Size())
	list.Show()
}
