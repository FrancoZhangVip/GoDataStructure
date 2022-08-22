package hot_one_hundred

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}
