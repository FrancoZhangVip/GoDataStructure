package leetcode_algorithm

import (
	"fmt"
	"testing"
)

func ShowTreeMidNode(node *TreeNode) {
	if node.Left != nil {
		ShowTreeMidNode(node.Left)
	}
	fmt.Print(node.Val, "\t")
	if node.Right != nil {
		ShowTreeMidNode(node.Right)
	}
}

func TestMergeTrees(t *testing.T) {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)
	node4 := new(TreeNode)

	node1.Val = 1
	node2.Val = 3
	node3.Val = 2
	node4.Val = 5

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4

	ShowTreeMidNode(node1)
	fmt.Println()
	///////////////////////

	node5 := new(TreeNode)
	node6 := new(TreeNode)
	node7 := new(TreeNode)
	node8 := new(TreeNode)
	node9 := new(TreeNode)

	node5.Val = 2
	node6.Val = 1
	node7.Val = 3
	node8.Val = 4
	node9.Val = 7

	node5.Left = node6
	node5.Right = node7
	node6.Right = node8
	node7.Right = node9

	ShowTreeMidNode(node5)
	fmt.Println()

	root := MergeTrees(node1, node5)
	ShowTreeMidNode(root)
	fmt.Println()
}
