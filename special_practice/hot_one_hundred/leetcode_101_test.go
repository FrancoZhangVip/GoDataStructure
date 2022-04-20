package hot_one_hundred

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)
	node4 := new(TreeNode)
	node5 := new(TreeNode)
	node6 := new(TreeNode)
	node7 := new(TreeNode)

	node1.Val = 1
	node2.Val = 2
	node3.Val = 2
	node4.Val = 3
	node5.Val = 4
	node6.Val = 4
	node7.Val = 3

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	fmt.Println(IsSymmetric(node1))

}

func TestIsSymmetric2(t *testing.T) {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)
	node4 := new(TreeNode)
	node5 := new(TreeNode)
	//node6 := new(TreeNode)
	node7 := new(TreeNode)

	node1.Val = 1
	node2.Val = 2
	node3.Val = 2
	node4.Val = 3
	node5.Val = 4
	//node6.Val = 4
	node7.Val = 3

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	//node3.Left = node6
	node3.Right = node7

	fmt.Println(IsSymmetric(node1))

}
