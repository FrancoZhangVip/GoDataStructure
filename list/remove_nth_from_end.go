package list

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	hiden := new(ListNode)
	hiden.Next = head
	i := 0
	quick := hiden
	slow := hiden
	for ; i < n; i++ {
		quick = quick.Next
	}
	for quick.Next != nil {
		quick = quick.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return hiden.Next
}
