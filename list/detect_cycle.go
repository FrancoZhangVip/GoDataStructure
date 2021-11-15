package list

/*
给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。
注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。


链接：https://leetcode-cn.com/problems/linked-list-cycle-ii

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	quick := head

	for true {
		if quick.Next == nil || quick.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		quick = quick.Next.Next
		if quick == slow {
			break
		}
	}
	index := head
	for true {
		if slow == index {
			return index
		}
		index = index.Next
		slow = slow.Next
	}
	return nil
}
