package algorithm_init

import (
	"github.com/hedzr/assert"
	"strconv"
	"testing"
)

/*
示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]
*/
func TestReverseList(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4
	node5 := new(ListNode)
	node5.Val = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	head := ReverseList(node1)
	answer1 := ""
	for head != nil {
		answer1 = answer1 + strconv.Itoa(head.Val) + "->"
		head = head.Next
	}
	answer1 = answer1 + "nil"
	println(answer1)
	assert.Equal(t, "5->4->3->2->1->nil", answer1)

	node1 = new(ListNode)
	node1.Val = 1
	node2 = new(ListNode)
	node2.Val = 2

	node1.Next = node2

	head = ReverseList(node1)
	answer2 := ""
	for head != nil {
		answer2 = answer2 + strconv.Itoa(head.Val) + "->"
		head = head.Next
	}
	answer2 = answer2 + "nil"
	println(answer2)
	assert.Equal(t, "2->1->nil", answer2)

	head = ReverseList(nil)
	answer3 := ""
	for head != nil {
		answer3 = answer3 + strconv.Itoa(head.Val) + "->"
		head = head.Next
	}
	answer3 = answer3 + "nil"
	println(answer3)
	assert.Equal(t, "nil", answer3)

}
