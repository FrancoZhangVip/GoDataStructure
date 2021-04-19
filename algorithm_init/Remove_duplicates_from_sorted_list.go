package algorithm_init

/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
返回同样按升序排列的结果链表。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	start := head
	travel := head.Next

	for travel != nil {
		if travel.Val != start.Val {
			start.Next = travel
			start = travel
			travel = travel.Next
		} else {
			travel = travel.Next
		}
	}

	start.Next = nil

	return head
}
