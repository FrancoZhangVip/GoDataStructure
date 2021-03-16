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

//先序遍历
func ShowTreeFrontNode(node *BSTNode) {
	fmt.Print(node.value, "\t")
	if node.lChild != nil {
		ShowTreeFrontNode(node.lChild)
	}
	if node.rChild != nil {
		ShowTreeFrontNode(node.rChild)
	}
}

func (tree *BinarySortedTree) ShowFront() {
	fmt.Println("先序遍历")
	if tree.root != nil {
		ShowTreeFrontNode(tree.root)
		fmt.Println()
	}
}

//中序遍历
func ShowTreeMidNode(node *BSTNode) {
	if node.lChild != nil {
		ShowTreeMidNode(node.lChild)
	}
	fmt.Print(node.value, "\t")
	if node.rChild != nil {
		ShowTreeMidNode(node.rChild)
	}
}

func (tree *BinarySortedTree) ShowMid() {
	fmt.Println("中序遍历")
	if tree.root != nil {
		ShowTreeMidNode(tree.root)
		fmt.Println()
	}
}

//后序遍历
func ShowTreeBackNode(node *BSTNode) {

	if node.lChild != nil {
		ShowTreeBackNode(node.lChild)
	}
	if node.rChild != nil {
		ShowTreeBackNode(node.rChild)
	}
	fmt.Print(node.value, "\t")
}

func (tree *BinarySortedTree) ShowBack() {
	fmt.Println("后序遍历")
	if tree.root != nil {
		ShowTreeBackNode(tree.root)
		fmt.Println()
	}
}
