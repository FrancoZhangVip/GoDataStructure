package list

/*
题意：删除链表中等于给定值 val 的所有节点。

示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例 2：
输入：head = [], val = 1
输出：[]

示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]

https://leetcode-cn.com/problems/remove-linked-list-elements/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	hidenHead := new(ListNode)
	hidenHead.Next = head

	slow := hidenHead
	fast := head

	for fast != nil {
		if fast.Val == val {
			fast = fast.Next
			slow.Next = fast
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}

	return hidenHead.Next
}

/*
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    head.Next = removeElements(head.Next, val)
    if head.Val == val {
        return head.Next
    }
    return head
}

*/
