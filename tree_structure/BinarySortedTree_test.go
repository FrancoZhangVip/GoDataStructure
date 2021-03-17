package tree_structure

import (
	"fmt"
	"testing"
)

func TestBinarySortedTree_Add(t *testing.T) {
	node1 := NewBSTNode(10)
	node2 := NewBSTNode(4)
	node3 := NewBSTNode(5)
	node4 := NewBSTNode(2)
	node5 := NewBSTNode(12)
	node6 := NewBSTNode(14)
	node7 := NewBSTNode(8)

	tree := NewBinarySortedTree()

	tree.Add(node1)
	tree.Add(node2)
	tree.Add(node3)
	tree.Add(node4)
	tree.Add(node5)
	tree.Add(node6)
	tree.Add(node7)

	tree.ShowFront()
	tree.ShowMid()
	tree.ShowBack()
}

func TestBinarySortedTree_Search(t *testing.T) {
	node1 := NewBSTNode(10)
	node2 := NewBSTNode(4)
	node3 := NewBSTNode(5)
	node4 := NewBSTNode(2)
	node5 := NewBSTNode(12)
	node6 := NewBSTNode(14)
	node7 := NewBSTNode(8)

	tree := NewBinarySortedTree()

	tree.Add(node1)
	tree.Add(node2)
	tree.Add(node3)
	tree.Add(node4)
	tree.Add(node5)
	tree.Add(node6)
	tree.Add(node7)

	tree.ShowMid()
	fmt.Println("Search: ", tree.Search(12))
	fmt.Println("Search: ", tree.Search(15))
}
