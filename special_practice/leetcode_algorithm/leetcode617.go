package leetcode_algorithm

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

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	root := new(TreeNode)
	if root1 != nil && root2 != nil {
		root.Val = root1.Val + root2.Val
		root.Left = MergeTrees(root1.Left, root2.Left)
		root.Right = MergeTrees(root1.Right, root2.Right)
	}
	if root1 == nil && root2 != nil {
		root.Val = root2.Val
		root.Left = MergeTrees(nil, root2.Left)
		root.Right = MergeTrees(nil, root2.Right)
	}
	if root1 != nil && root2 == nil {
		root.Val = root1.Val
		root.Left = MergeTrees(root1.Left, nil)
		root.Right = MergeTrees(root1.Right, nil)
	}

	return root
}
