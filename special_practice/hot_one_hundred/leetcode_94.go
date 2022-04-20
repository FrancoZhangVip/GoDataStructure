package hot_one_hundred

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	return MidOrder(root, res)
}

func MidOrder(root *TreeNode, res []int) []int {
	if root.Left != nil {
		res = MidOrder(root.Left, res)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = MidOrder(root.Right, res)
	}
	return res
}
