package hot_one_hundred

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left != nil && root.Right != nil && root.Left.Val == root.Right.Val {
		return CheckMirror(root.Left, root.Right)
	} else {
		return false
	}
}

func CheckMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil && root1.Val == root2.Val {
		if CheckMirror(root1.Left, root2.Right) && CheckMirror(root1.Right, root2.Left) {
			return true
		}
		return false
	}
	return false
}
