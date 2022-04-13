package leetcode_algorithm

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

func MiddleNode(head *ListNode) *ListNode {
	p := head
	q := head.Next
	for true {
		if q == nil {
			return p
		}
		p = p.Next
		q = q.Next
		if q == nil {
			return p
		}
		q = q.Next
	}
	return nil
}
