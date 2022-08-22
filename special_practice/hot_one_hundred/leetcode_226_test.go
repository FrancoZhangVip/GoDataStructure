package hot_one_hundred

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)
	node4 := new(TreeNode)
	node5 := new(TreeNode)
	node6 := new(TreeNode)
	node7 := new(TreeNode)

	node1.Val = 4
	node2.Val = 2
	node3.Val = 7
	node4.Val = 1
	node5.Val = 3
	node6.Val = 6
	node7.Val = 9

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node4.Right = node5
	node3.Left = node6
	node3.Right = node7

	InvertTree(node1)

	ShowTree(node1)

	println(" ", node1.Right.Val)
}

func ShowTree(root *TreeNode) {
	if root == nil {
		return
	}
	ShowTree(root.Left)
	fmt.Print(root.Val, "\t")
	ShowTree(root.Right)
}
