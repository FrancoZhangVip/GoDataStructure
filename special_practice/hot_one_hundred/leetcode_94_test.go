package hot_one_hundred

import "testing"

func TestInorderTraversal(t *testing.T) {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)

	node1.Val = 1
	node2.Val = 2
	node3.Val = 3

	node1.Right = node2
	node2.Left = node3

	res := InorderTraversal(node1)
	ShowArray(res)
}
