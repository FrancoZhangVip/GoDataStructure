package hot_one_hundred

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	return SearchDepth(root, depth)
}

func SearchDepth(root *TreeNode, depth int) int {
	leftdepth, rightdepth, tempdepth := 0, 0, depth
	if root.Left != nil {
		leftdepth = SearchDepth(root.Left, depth+1)
	}
	if root.Right != nil {
		rightdepth = SearchDepth(root.Right, depth+1)
	}
	if leftdepth > tempdepth {
		tempdepth = leftdepth
	}
	if rightdepth > tempdepth {
		tempdepth = rightdepth
	}
	return tempdepth
}
