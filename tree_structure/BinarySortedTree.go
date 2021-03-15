package tree_structure

import "fmt"

type BSTNode struct {
	value  int
	lChild *BSTNode
	rChild *BSTNode
}

type BinarySortedTree struct {
	root *BSTNode
}

func NewBSTNode(val int) *BSTNode {
	node := new(BSTNode)
	node.value = val
	return node
}

func NewBinarySortedTree() *BinarySortedTree {
	return new(BinarySortedTree)
}

func SearchIndex(p *BSTNode, node *BSTNode) {
	if p.lChild == nil && p.value > node.value {
		p.lChild = node
		return
	}

	if p.rChild == nil && p.value <= node.value {
		p.rChild = node
		return
	}

	if p.lChild != nil && p.value > node.value {
		SearchIndex(p.lChild, node)
	}

	if p.rChild != nil && p.value <= node.value {
		SearchIndex(p.rChild, node)
	}
}

func (tree *BinarySortedTree) Add(node *BSTNode) {
	if tree.root == nil {
		tree.root = node
		return
	}
	SearchIndex(tree.root, node)
}

func ShowTreeNode(node *BSTNode) {
	if node.lChild != nil {
		ShowTreeNode(node.lChild)
	}
	fmt.Print(node.value, "\t")
	if node.rChild != nil {
		ShowTreeNode(node.rChild)
	}
}

func (tree *BinarySortedTree) Show() {
	if tree.root != nil {
		ShowTreeNode(tree.root)
		fmt.Println()
	}
}
