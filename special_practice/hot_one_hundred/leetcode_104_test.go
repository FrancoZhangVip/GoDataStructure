package hot_one_hundred

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)
	node4 := new(TreeNode)
	node5 := new(TreeNode)

	node1.Val = 3
	node2.Val = 9
	node3.Val = 20
	node4.Val = 15
	node5.Val = 7

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	assert.Equal(t, 3, MaxDepth(node1))
}
